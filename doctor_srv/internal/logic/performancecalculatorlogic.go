package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/qas491/hospital/doctor_srv/internal/svc"
	"github.com/qas491/hospital/doctor_srv/model/mysql"

	"github.com/zeromicro/go-zero/core/logx"
)

// PerformanceCalculatorLogic 业绩计算逻辑处理器
type PerformanceCalculatorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewPerformanceCalculatorLogic 创建业绩计算逻辑处理器实例
func NewPerformanceCalculatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformanceCalculatorLogic {
	return &PerformanceCalculatorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CalculatePerformance 计算医生业绩
// @param paymentID 缴费ID
// @return error 错误信息
func (l *PerformanceCalculatorLogic) CalculatePerformance(paymentID string) error {
	// 1. 查询缴费信息
	var payment mysql.HisPatientPayment
	if err := l.svcCtx.DB.Where("payment_id = ? AND payment_status = 'success'", paymentID).First(&payment).Error; err != nil {
		l.Logger.Errorf("查询缴费信息失败: %v", err)
		return fmt.Errorf("查询缴费信息失败: %v", err)
	}

	// 2. 查询处方信息
	var careOrder mysql.HisCareOrder
	if err := l.svcCtx.DB.Where("co_id = ?", payment.CoID).First(&careOrder).Error; err != nil {
		l.Logger.Errorf("查询处方信息失败: %v", err)
		return fmt.Errorf("查询处方信息失败: %v", err)
	}

	// 3. 查询适用的业绩规则
	rule, err := l.getPerformanceRule(careOrder.CoType, careOrder.UserID)
	if err != nil {
		l.Logger.Errorf("查询业绩规则失败: %v", err)
		return fmt.Errorf("查询业绩规则失败: %v", err)
	}

	// 4. 计算业绩金额
	performanceAmount := l.calculatePerformanceAmount(payment.PaymentAmount, rule)

	// 5. 生成业绩明细记录
	performanceDetail := &mysql.HisPerformanceDetail{
		PerformanceID:     l.generatePerformanceID(),
		CoID:              payment.CoID,
		PaymentID:         payment.PaymentID,
		DoctorID:          payment.DoctorID,
		DoctorName:        payment.DoctorName,
		DeptID:            careOrder.UserID, // 这里需要根据实际情况获取科室ID
		DeptName:          "",               // 需要根据科室ID查询科室名称
		PaymentAmount:     payment.PaymentAmount,
		PerformanceRate:   rule.RuleValue,
		PerformanceAmount: performanceAmount,
		RuleID:            rule.RuleID,
		PerformanceDate:   &time.Time{},
		WeekStart:         l.getWeekStart(),
		WeekEnd:           l.getWeekEnd(),
		CreateTime:        &time.Time{},
		UpdateTime:        &time.Time{},
		CreateBy:          "system",
		UpdateBy:          "system",
	}

	// 6. 保存业绩明细
	if err := l.svcCtx.DB.Create(performanceDetail).Error; err != nil {
		l.Logger.Errorf("保存业绩明细失败: %v", err)
		return fmt.Errorf("保存业绩明细失败: %v", err)
	}

	l.Logger.Infof("成功计算医生 %s 的业绩: %.2f", payment.DoctorName, performanceAmount)
	return nil
}

// getPerformanceRule 获取适用的业绩规则
func (l *PerformanceCalculatorLogic) getPerformanceRule(coType string, doctorID int64) (*mysql.HisPerformanceRule, error) {
	var rule mysql.HisPerformanceRule

	// 查询当前生效的业绩规则
	now := time.Now()
	err := l.svcCtx.DB.Where(`
		(co_type = ? OR co_type = '') 
		AND status = 'active' 
		AND (effective_date IS NULL OR effective_date <= ?) 
		AND (expiry_date IS NULL OR expiry_date >= ?)
		ORDER BY rule_value DESC
		LIMIT 1
	`, coType, now, now).First(&rule).Error

	if err != nil {
		// 如果没有找到特定规则，使用默认规则
		err = l.svcCtx.DB.Where(`
			status = 'active' 
			AND (effective_date IS NULL OR effective_date <= ?) 
			AND (expiry_date IS NULL OR expiry_date >= ?)
			ORDER BY rule_value DESC
			LIMIT 1
		`, now, now).First(&rule).Error
	}

	return &rule, err
}

// calculatePerformanceAmount 计算业绩金额
func (l *PerformanceCalculatorLogic) calculatePerformanceAmount(paymentAmount float64, rule *mysql.HisPerformanceRule) float64 {
	var performanceAmount float64

	switch rule.RuleType {
	case "percentage":
		// 百分比计算
		performanceAmount = paymentAmount * (rule.RuleValue / 100.0)
	case "fixed":
		// 固定金额
		performanceAmount = rule.RuleValue
	default:
		// 默认百分比计算
		performanceAmount = paymentAmount * (rule.RuleValue / 100.0)
	}

	// 应用最小/最大金额限制
	if rule.MinAmount > 0 && performanceAmount < rule.MinAmount {
		performanceAmount = rule.MinAmount
	}
	if rule.MaxAmount > 0 && performanceAmount > rule.MaxAmount {
		performanceAmount = rule.MaxAmount
	}

	return performanceAmount
}

// generatePerformanceID 生成业绩ID
func (l *PerformanceCalculatorLogic) generatePerformanceID() string {
	return fmt.Sprintf("PERF_%d", time.Now().UnixNano())
}

// getWeekStart 获取本周开始时间
func (l *PerformanceCalculatorLogic) getWeekStart() *time.Time {
	now := time.Now()
	weekday := now.Weekday()
	if weekday == time.Sunday {
		weekday = 7
	} else {
		weekday--
	}
	weekStart := now.AddDate(0, 0, -int(weekday))
	weekStart = time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, weekStart.Location())
	return &weekStart
}

// getWeekEnd 获取本周结束时间
func (l *PerformanceCalculatorLogic) getWeekEnd() *time.Time {
	weekStart := l.getWeekStart()
	weekEnd := weekStart.AddDate(0, 0, 6)
	weekEnd = time.Date(weekEnd.Year(), weekEnd.Month(), weekEnd.Day(), 23, 59, 59, 999999999, weekEnd.Location())
	return &weekEnd
}
