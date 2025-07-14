package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/qas491/hospital/doctor_srv/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// GetDoctorPerformanceLogic 获取医生业绩逻辑结构体
type GetDoctorPerformanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewGetDoctorPerformanceLogic 创建获取医生业绩逻辑实例
func NewGetDoctorPerformanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDoctorPerformanceLogic {
	return &GetDoctorPerformanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetDoctorPerformance 获取医生业绩主方法
// 根据医生ID和日期范围查询医生的业绩统计信息和明细数据
func (l *GetDoctorPerformanceLogic) GetDoctorPerformance(in *doctor.GetDoctorPerformanceReq) (*doctor.GetDoctorPerformanceResp, error) {
	// 解析日期参数
	startDate := in.StartDate
	endDate := in.EndDate

	// 如果没有指定日期，默认查询本月
	if startDate == "" || endDate == "" {
		now := time.Now()
		startDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	}

	// 查询医生基本信息
	doctorInfo, err := l.getDoctorInfo(in.DoctorId)
	if err != nil {
		l.Error("查询医生信息失败", logx.Field("error", err), logx.Field("doctor_id", in.DoctorId))
		return &doctor.GetDoctorPerformanceResp{
			Code:    500,
			Message: "查询医生信息失败",
		}, nil
	}

	// 查询医生业绩数据
	performanceData, err := l.getDoctorPerformanceData(in.DoctorId, startDate, endDate)
	if err != nil {
		l.Error("查询医生业绩失败", logx.Field("error", err), logx.Field("doctor_id", in.DoctorId))
		return &doctor.GetDoctorPerformanceResp{
			Code:    500,
			Message: "查询医生业绩失败",
		}, nil
	}

	// 查询业绩明细
	performanceDetails, err := l.getPerformanceDetails(in.DoctorId, startDate, endDate)
	if err != nil {
		l.Error("查询业绩明细失败", logx.Field("error", err), logx.Field("doctor_id", in.DoctorId))
		return &doctor.GetDoctorPerformanceResp{
			Code:    500,
			Message: "查询业绩明细失败",
		}, nil
	}

	l.Info("获取医生业绩成功",
		logx.Field("doctor_id", in.DoctorId),
		logx.Field("start_date", startDate),
		logx.Field("end_date", endDate),
		logx.Field("total_performance", performanceData.TotalPerformance),
		logx.Field("prescription_count", performanceData.PrescriptionCount))

	return &doctor.GetDoctorPerformanceResp{
		Code:    200,
		Message: "获取成功",
		Performance: &doctor.DoctorPerformanceInfo{
			DoctorId:           in.DoctorId,
			DoctorName:         doctorInfo.DoctorName,
			DeptName:           doctorInfo.DeptName,
			TotalPerformance:   performanceData.TotalPerformance,
			PrescriptionCount:  performanceData.PrescriptionCount,
			PaymentCount:       performanceData.PaymentCount,
			PerformanceDetails: performanceDetails,
		},
	}, nil
}

// getDoctorInfo 查询医生基本信息
// 根据医生ID查询医生的姓名和所属科室信息
func (l *GetDoctorPerformanceLogic) getDoctorInfo(doctorID int64) (*DoctorInfo, error) {
	var doctorInfo DoctorInfo

	query := `
		SELECT 
			d.user_id as doctor_id,
			d.user_name as doctor_name,
			dept.dept_name
		FROM sys_user d
		LEFT JOIN sys_dept dept ON d.dept_id = dept.dept_id
		WHERE d.user_id = ? AND d.user_type = 'doctor' AND d.del_flag = '0'
	`

	err := l.svcCtx.DB.Raw(query, doctorID).Scan(&doctorInfo).Error
	if err != nil {
		return nil, fmt.Errorf("查询医生信息失败: %v", err)
	}

	if doctorInfo.DoctorID == 0 {
		return nil, fmt.Errorf("医生不存在: %d", doctorID)
	}

	return &doctorInfo, nil
}

// getDoctorPerformanceData 查询医生业绩数据
// 统计指定时间范围内医生的处方数量、总业绩和已支付数量
func (l *GetDoctorPerformanceLogic) getDoctorPerformanceData(doctorID int64, startDate, endDate string) (*PerformanceData, error) {
	var performanceData PerformanceData

	query := `
		SELECT 
			COUNT(DISTINCT p.co_id) as prescription_count,
			COALESCE(SUM(p.all_amount), 0) as total_performance,
			COUNT(DISTINCT CASE WHEN p.status = 'paid' THEN p.co_id END) as payment_count
		FROM prescription p
		WHERE p.user_id = ? 
			AND DATE(p.create_time) >= ? 
			AND DATE(p.create_time) <= ?
	`

	err := l.svcCtx.DB.Raw(query, doctorID, startDate, endDate).Scan(&performanceData).Error
	if err != nil {
		return nil, fmt.Errorf("查询医生业绩数据失败: %v", err)
	}

	return &performanceData, nil
}

// getPerformanceDetails 查询业绩明细
// 查询指定时间范围内医生的所有已支付处方明细
func (l *GetDoctorPerformanceLogic) getPerformanceDetails(doctorID int64, startDate, endDate string) ([]*doctor.PerformanceDetail, error) {
	var details []PerformanceDetailItem

	query := `
		SELECT 
			CONCAT('PERF_', p.co_id) as performance_id,
			p.co_id,
			p.all_amount as payment_amount,
			p.all_amount as performance_amount,
			1.0 as performance_rate,
			DATE(p.create_time) as performance_date
		FROM prescription p
		WHERE p.user_id = ? 
			AND DATE(p.create_time) >= ? 
			AND DATE(p.create_time) <= ?
			AND p.status = 'paid'
		ORDER BY p.create_time DESC
	`

	err := l.svcCtx.DB.Raw(query, doctorID, startDate, endDate).Scan(&details).Error
	if err != nil {
		return nil, fmt.Errorf("查询业绩明细失败: %v", err)
	}

	// 转换为proto格式
	var performanceDetails []*doctor.PerformanceDetail
	for _, detail := range details {
		performanceDetails = append(performanceDetails, &doctor.PerformanceDetail{
			PerformanceId:     detail.PerformanceID,
			CoId:              detail.CoID,
			PaymentAmount:     detail.PaymentAmount,
			PerformanceAmount: detail.PerformanceAmount,
			PerformanceRate:   detail.PerformanceRate,
			PerformanceDate:   detail.PerformanceDate,
		})
	}

	return performanceDetails, nil
}

// DoctorInfo 医生信息结构
// 用于存储医生的基本信息
type DoctorInfo struct {
	DoctorID   int64  `json:"doctor_id" gorm:"column:doctor_id"`     // 医生ID
	DoctorName string `json:"doctor_name" gorm:"column:doctor_name"` // 医生姓名
	DeptName   string `json:"dept_name" gorm:"column:dept_name"`     // 科室名称
}

// PerformanceData 业绩数据结构
// 用于存储医生的业绩统计信息
type PerformanceData struct {
	PrescriptionCount int64   `json:"prescription_count" gorm:"column:prescription_count"` // 处方数量
	TotalPerformance  float64 `json:"total_performance" gorm:"column:total_performance"`   // 总业绩
	PaymentCount      int64   `json:"payment_count" gorm:"column:payment_count"`           // 已支付数量
}

// PerformanceDetailItem 业绩明细项结构
// 用于存储单条业绩明细信息
type PerformanceDetailItem struct {
	PerformanceID     string  `json:"performance_id" gorm:"column:performance_id"`         // 业绩ID
	CoID              string  `json:"co_id" gorm:"column:co_id"`                           // 处方编号
	PaymentAmount     float64 `json:"payment_amount" gorm:"column:payment_amount"`         // 支付金额
	PerformanceAmount float64 `json:"performance_amount" gorm:"column:performance_amount"` // 业绩金额
	PerformanceRate   float64 `json:"performance_rate" gorm:"column:performance_rate"`     // 业绩比例
	PerformanceDate   string  `json:"performance_date" gorm:"column:performance_date"`     // 业绩日期
}
