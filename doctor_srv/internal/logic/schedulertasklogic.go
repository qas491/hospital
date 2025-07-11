package logic

import (
	"context"
	"time"

	"github.com/qas491/hospital/doctor_srv/internal/svc"
	"github.com/qas491/hospital/doctor_srv/model/mysql"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

// SchedulerTaskLogic 定时任务逻辑处理器
type SchedulerTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewSchedulerTaskLogic 创建定时任务逻辑处理器实例
func NewSchedulerTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SchedulerTaskLogic {
	return &SchedulerTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// StartWeeklyRankingScheduler 启动周排行榜定时任务
// 每周一凌晨2点执行
func (l *SchedulerTaskLogic) StartWeeklyRankingScheduler() {
	// 计算下次执行时间（下周一凌晨2点）
	nextRun := l.getNextMonday2AM()

	l.Logger.Infof("周排行榜定时任务已启动，下次执行时间: %s", nextRun.Format("2006-01-02 15:04:05"))

	// 启动定时任务
	threading.GoSafe(func() {
		l.runWeeklyRankingScheduler(nextRun)
	})
}

// runWeeklyRankingScheduler 运行周排行榜定时任务
func (l *SchedulerTaskLogic) runWeeklyRankingScheduler(nextRun time.Time) {
	for {
		now := time.Now()

		// 如果到达执行时间
		if now.After(nextRun) {
			l.Logger.Info("开始执行周排行榜生成任务")

			// 执行排行榜生成
			rankingLogic := NewWeeklyRankingLogic(l.ctx, l.svcCtx)
			if err := rankingLogic.GenerateWeeklyRanking(); err != nil {
				l.Logger.Errorf("生成周排行榜失败: %v", err)
			} else {
				l.Logger.Info("周排行榜生成完成")
			}

			// 计算下次执行时间
			nextRun = l.getNextMonday2AM()
			l.Logger.Infof("下次执行时间: %s", nextRun.Format("2006-01-02 15:04:05"))
		}

		// 休眠1分钟检查一次
		time.Sleep(1 * time.Minute)
	}
}

// getNextMonday2AM 获取下周一凌晨2点的时间
func (l *SchedulerTaskLogic) getNextMonday2AM() time.Time {
	now := time.Now()

	// 计算到下周一的天数
	daysUntilMonday := (8 - int(now.Weekday())) % 7
	if daysUntilMonday == 0 {
		daysUntilMonday = 7
	}

	// 计算下周一凌晨2点
	nextMonday := now.AddDate(0, 0, daysUntilMonday)
	nextMonday2AM := time.Date(
		nextMonday.Year(),
		nextMonday.Month(),
		nextMonday.Day(),
		2, 0, 0, 0,
		nextMonday.Location(),
	)

	return nextMonday2AM
}

// StartPerformanceCalculationScheduler 启动业绩计算定时任务
// 每5分钟检查一次缴费成功的记录并计算业绩
func (l *SchedulerTaskLogic) StartPerformanceCalculationScheduler() {
	l.Logger.Info("业绩计算定时任务已启动")

	threading.GoSafe(func() {
		l.runPerformanceCalculationScheduler()
	})
}

// runPerformanceCalculationScheduler 运行业绩计算定时任务
func (l *SchedulerTaskLogic) runPerformanceCalculationScheduler() {
	for {
		// 查询未计算业绩的缴费记录
		var payments []mysql.HisPatientPayment
		err := l.svcCtx.DB.Where(`
			payment_status = 'success' 
			AND payment_id NOT IN (
				SELECT DISTINCT payment_id FROM his_performance_detail
			)
		`).Find(&payments).Error

		if err != nil {
			l.Logger.Errorf("查询未计算业绩的缴费记录失败: %v", err)
		} else if len(payments) > 0 {
			l.Logger.Infof("发现 %d 条未计算业绩的缴费记录", len(payments))

			// 计算业绩
			performanceLogic := NewPerformanceCalculatorLogic(l.ctx, l.svcCtx)
			for _, payment := range payments {
				if err := performanceLogic.CalculatePerformance(payment.PaymentID); err != nil {
					l.Logger.Errorf("计算缴费记录 %s 的业绩失败: %v", payment.PaymentID, err)
				}
			}
		}

		// 休眠5分钟
		time.Sleep(5 * time.Minute)
	}
}
