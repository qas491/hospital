package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/qas491/hospital/doctor_srv/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// GenerateWeeklyRankingLogic 生成周排行榜逻辑结构体
type GenerateWeeklyRankingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewGenerateWeeklyRankingLogic 创建生成周排行榜逻辑实例
func NewGenerateWeeklyRankingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateWeeklyRankingLogic {
	return &GenerateWeeklyRankingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GenerateWeeklyRanking 生成周排行榜主方法
// 该方法会查询本周所有医生的业绩数据，计算排行榜并存储到Redis中
func (l *GenerateWeeklyRankingLogic) GenerateWeeklyRanking(in *doctor.GenerateWeeklyRankingReq) (*doctor.GenerateWeeklyRankingResp, error) {
	// 获取本周的开始和结束时间
	weekStart, weekEnd := l.getWeekStartEnd()

	// 从数据库查询本周所有医生的业绩数据
	doctorPerformances, err := l.getDoctorPerformancesFromDB(weekStart, weekEnd)
	if err != nil {
		l.Error("查询医生业绩失败", logx.Field("error", err))
		return &doctor.GenerateWeeklyRankingResp{
			Code:    500,
			Message: "查询医生业绩失败",
			Success: false,
		}, nil
	}

	// 计算排行榜并存储到Redis
	err = l.saveRankingToRedis(doctorPerformances)
	if err != nil {
		l.Error("保存排行榜到Redis失败", logx.Field("error", err))
		return &doctor.GenerateWeeklyRankingResp{
			Code:    500,
			Message: "保存排行榜失败",
			Success: false,
		}, nil
	}

	l.Info("周排行榜生成成功",
		logx.Field("doctor_count", len(doctorPerformances)),
		logx.Field("week_start", weekStart.Format("2006-01-02")),
		logx.Field("week_end", weekEnd.Format("2006-01-02")))

	return &doctor.GenerateWeeklyRankingResp{
		Code:        200,
		Message:     "周排行榜生成成功",
		Success:     true,
		DoctorCount: int64(len(doctorPerformances)),
	}, nil
}

// getWeekStartEnd 获取本周的开始和结束时间
// 返回本周一00:00:00到周日23:59:59的时间范围
func (l *GenerateWeeklyRankingLogic) getWeekStartEnd() (time.Time, time.Time) {
	now := time.Now()
	weekday := now.Weekday()
	if weekday == time.Sunday {
		weekday = 7
	} else {
		weekday--
	}
	weekStart := now.AddDate(0, 0, -int(weekday))
	weekStart = time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, weekStart.Location())
	weekEnd := weekStart.AddDate(0, 0, 6)
	weekEnd = time.Date(weekEnd.Year(), weekEnd.Month(), weekEnd.Day(), 23, 59, 59, 999999999, weekEnd.Location())
	return weekStart, weekEnd
}

// getDoctorPerformancesFromDB 从数据库查询医生业绩数据
// 查询指定时间范围内所有医生的处方数量和总业绩
func (l *GenerateWeeklyRankingLogic) getDoctorPerformancesFromDB(startTime, endTime time.Time) ([]DoctorPerformance, error) {
	var performances []DoctorPerformance

	// 查询本周所有医生的处方和业绩数据
	query := `
		SELECT 
			d.user_id as doctor_id,
			d.user_name as doctor_name,
			dept.dept_name,
			COUNT(DISTINCT p.co_id) as prescription_count,
			COALESCE(SUM(p.all_amount), 0) as total_performance
		FROM sys_user d
		LEFT JOIN sys_dept dept ON d.dept_id = dept.dept_id
		LEFT JOIN prescription p ON d.user_id = p.user_id 
			AND p.create_time >= ? AND p.create_time <= ?
		WHERE d.user_type = 'doctor' AND d.del_flag = '0'
		GROUP BY d.user_id, d.user_name, dept.dept_name
		HAVING total_performance > 0
		ORDER BY total_performance DESC
	`

	err := l.svcCtx.DB.Raw(query, startTime, endTime).Scan(&performances).Error
	if err != nil {
		return nil, fmt.Errorf("查询医生业绩失败: %v", err)
	}

	return performances, nil
}

// saveRankingToRedis 保存排行榜到Redis
// 使用Redis的有序集合存储排行榜数据，以业绩作为分数进行排序
func (l *GenerateWeeklyRankingLogic) saveRankingToRedis(performances []DoctorPerformance) error {
	redisKey := "weekly_ranking"

	// 删除旧的排行榜数据
	err := l.svcCtx.RedisClient.Del(context.Background(), redisKey).Err()
	if err != nil {
		return fmt.Errorf("删除旧排行榜失败: %v", err)
	}

	// 添加新的排行榜数据
	for _, performance := range performances {
		// 构造Redis成员数据：doctor_id:doctor_name:dept_name:prescription_count
		member := fmt.Sprintf("%d:%s:%s:%d",
			performance.DoctorID,
			performance.DoctorName,
			performance.DeptName,
			performance.PrescriptionCount)

		// 使用总业绩作为分数，按降序排列
		score := performance.TotalPerformance

		err := l.svcCtx.RedisClient.ZAdd(context.Background(), redisKey, &redis.Z{
			Score:  score,
			Member: member,
		}).Err()

		if err != nil {
			return fmt.Errorf("添加排行榜数据失败: %v", err)
		}
	}

	return nil
}

// DoctorPerformance 医生业绩数据结构
// 用于存储从数据库查询到的医生业绩信息
type DoctorPerformance struct {
	DoctorID          int64   `json:"doctor_id" gorm:"column:doctor_id"`                   // 医生ID
	DoctorName        string  `json:"doctor_name" gorm:"column:doctor_name"`               // 医生姓名
	DeptName          string  `json:"dept_name" gorm:"column:dept_name"`                   // 科室名称
	PrescriptionCount int64   `json:"prescription_count" gorm:"column:prescription_count"` // 处方数量
	TotalPerformance  float64 `json:"total_performance" gorm:"column:total_performance"`   // 总业绩
}
