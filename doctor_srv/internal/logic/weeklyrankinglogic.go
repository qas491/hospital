package logic

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	redisv8 "github.com/go-redis/redis/v8"
	"github.com/qas491/hospital/doctor_srv/internal/svc"
	"github.com/qas491/hospital/doctor_srv/model/mysql"

	"github.com/zeromicro/go-zero/core/logx"
)

// WeeklyRankingLogic 周排行榜逻辑处理器
type WeeklyRankingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewWeeklyRankingLogic 创建周排行榜逻辑处理器实例
func NewWeeklyRankingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WeeklyRankingLogic {
	return &WeeklyRankingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GenerateWeeklyRanking 生成周排行榜
// @return error 错误信息
func (l *WeeklyRankingLogic) GenerateWeeklyRanking() error {
	weekStart := l.getWeekStart()
	weekEnd := l.getWeekEnd()

	l.Logger.Infof("开始生成周排行榜，时间范围: %s 到 %s", weekStart.Format("2006-01-02"), weekEnd.Format("2006-01-02"))

	// 1. 查询本周所有医生的业绩汇总
	var performanceSummary []struct {
		DoctorID          int64   `json:"doctor_id"`
		DoctorName        string  `json:"doctor_name"`
		DeptID            int64   `json:"dept_id"`
		DeptName          string  `json:"dept_name"`
		TotalPerformance  float64 `json:"total_performance"`
		PrescriptionCount int     `json:"prescription_count"`
	}

	err := l.svcCtx.DB.Table("his_performance_detail").
		Select(`
			doctor_id,
			doctor_name,
			dept_id,
			dept_name,
			SUM(performance_amount) as total_performance,
			COUNT(DISTINCT co_id) as prescription_count
		`).
		Where("performance_date >= ? AND performance_date <= ?", weekStart, weekEnd).
		Group("doctor_id, doctor_name, dept_id, dept_name").
		Order("total_performance DESC").
		Scan(&performanceSummary).Error

	if err != nil {
		l.Logger.Errorf("查询业绩汇总失败: %v", err)
		return fmt.Errorf("查询业绩汇总失败: %v", err)
	}

	// 2. 删除上周排行榜数据
	l.deleteLastWeekRanking()

	// 3. 生成新的排行榜数据
	var rankings []mysql.HisWeeklyRanking
	now := time.Now()

	for i, summary := range performanceSummary {
		ranking := mysql.HisWeeklyRanking{
			RankingID:         l.generateRankingID(),
			DoctorID:          summary.DoctorID,
			DoctorName:        summary.DoctorName,
			DeptID:            summary.DeptID,
			DeptName:          summary.DeptName,
			WeekStart:         &weekStart,
			WeekEnd:           &weekEnd,
			TotalPerformance:  summary.TotalPerformance,
			Rank:              i + 1,
			PrescriptionCount: summary.PrescriptionCount,
			CreateTime:        &now,
			UpdateTime:        &now,
		}
		rankings = append(rankings, ranking)
	}

	// 4. 保存排行榜到数据库
	if len(rankings) > 0 {
		if err := l.svcCtx.DB.Create(&rankings).Error; err != nil {
			l.Logger.Errorf("保存排行榜失败: %v", err)
			return fmt.Errorf("保存排行榜失败: %v", err)
		}
	}

	// 5. 更新Redis排行榜
	if err := l.updateRedisRanking(rankings); err != nil {
		l.Logger.Errorf("更新Redis排行榜失败: %v", err)
		return fmt.Errorf("更新Redis排行榜失败: %v", err)
	}

	l.Logger.Infof("成功生成周排行榜，共 %d 名医生", len(rankings))
	return nil
}

// deleteLastWeekRanking 删除上周排行榜数据
func (l *WeeklyRankingLogic) deleteLastWeekRanking() {
	// 删除数据库中的旧数据
	l.svcCtx.DB.Where("week_start < ?", l.getWeekStart()).Delete(&mysql.HisWeeklyRanking{})

	// 删除Redis中的旧数据
	redisKey := "weekly_ranking"
	l.svcCtx.RedisClient.Del(context.Background(), redisKey)
}

// updateRedisRanking 更新Redis排行榜
func (l *WeeklyRankingLogic) updateRedisRanking(rankings []mysql.HisWeeklyRanking) error {
	redisKey := "weekly_ranking"

	// 使用Redis有序集合存储排行榜
	for _, ranking := range rankings {
		// 使用业绩金额作为分数，医生ID作为成员
		score := ranking.TotalPerformance
		member := fmt.Sprintf("%d:%s:%s:%d",
			ranking.DoctorID,
			ranking.DoctorName,
			ranking.DeptName,
			ranking.PrescriptionCount)

		// 添加到有序集合
		err := l.svcCtx.RedisClient.ZAdd(context.Background(), redisKey, &redisv8.Z{
			Score:  score,
			Member: member,
		}).Err()
		if err != nil {
			return fmt.Errorf("添加排行榜数据到Redis失败: %v", err)
		}
	}

	// 设置过期时间（7天）
	l.svcCtx.RedisClient.Expire(context.Background(), redisKey, 7*24*time.Hour)

	return nil
}

// GetWeeklyRanking 获取周排行榜
// @param limit 获取前N名
// @return []map[string]interface{} 排行榜数据
// @return error 错误信息
func (l *WeeklyRankingLogic) GetWeeklyRanking(limit int) ([]map[string]interface{}, error) {
	redisKey := "weekly_ranking"

	// 从Redis获取排行榜数据（按分数降序）
	result, err := l.svcCtx.RedisClient.ZRevRangeWithScores(context.Background(), redisKey, 0, int64(limit-1)).Result()
	if err != nil {
		l.Logger.Errorf("从Redis获取排行榜失败: %v", err)
		return nil, fmt.Errorf("从Redis获取排行榜失败: %v", err)
	}

	var rankings []map[string]interface{}
	for i, item := range result {
		// 解析成员数据
		member := item.Member.(string)
		parts := strings.Split(member, ":")
		if len(parts) >= 4 {
			doctorID, _ := strconv.ParseInt(parts[0], 10, 64)
			doctorName := parts[1]
			deptName := parts[2]
			prescriptionCount, _ := strconv.Atoi(parts[3])

			ranking := map[string]interface{}{
				"rank":               i + 1,
				"doctor_id":          doctorID,
				"doctor_name":        doctorName,
				"dept_name":          deptName,
				"total_performance":  item.Score,
				"prescription_count": prescriptionCount,
			}
			rankings = append(rankings, ranking)
		}
	}

	return rankings, nil
}

// generateRankingID 生成排行榜ID
func (l *WeeklyRankingLogic) generateRankingID() string {
	return fmt.Sprintf("RANK_%d", time.Now().UnixNano())
}

// getWeekStart 获取本周开始时间
func (l *WeeklyRankingLogic) getWeekStart() time.Time {
	now := time.Now()
	weekday := now.Weekday()
	if weekday == time.Sunday {
		weekday = 7
	} else {
		weekday--
	}
	weekStart := now.AddDate(0, 0, -int(weekday))
	return time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, weekStart.Location())
}

// getWeekEnd 获取本周结束时间
func (l *WeeklyRankingLogic) getWeekEnd() time.Time {
	weekStart := l.getWeekStart()
	weekEnd := weekStart.AddDate(0, 0, 6)
	return time.Date(weekEnd.Year(), weekEnd.Month(), weekEnd.Day(), 23, 59, 59, 999999999, weekEnd.Location())
}
