package logic

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/qas491/hospital/doctor_srv/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// GetWeeklyRankingLogic 获取周排行榜逻辑结构体
type GetWeeklyRankingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewGetWeeklyRankingLogic 创建获取周排行榜逻辑实例
func NewGetWeeklyRankingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWeeklyRankingLogic {
	return &GetWeeklyRankingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetWeeklyRanking 获取周排行榜主方法
// 从Redis中获取本周医生业绩排行榜数据，支持限制返回数量
func (l *GetWeeklyRankingLogic) GetWeeklyRanking(in *doctor.GetWeeklyRankingReq) (*doctor.GetWeeklyRankingResp, error) {
	// 设置默认限制数量
	limit := in.Limit
	if limit <= 0 {
		limit = 10
	}

	// 从Redis获取排行榜数据
	rankings, err := l.getWeeklyRankingFromRedis(int(limit))
	if err != nil {
		l.Error("获取排行榜失败", logx.Field("error", err))
		return &doctor.GetWeeklyRankingResp{
			Code:    500,
			Message: "获取排行榜失败",
		}, nil
	}

	// 获取本周的开始和结束时间
	weekStart, weekEnd := l.getWeekStartEnd()

	l.Info("获取周排行榜成功",
		logx.Field("limit", limit),
		logx.Field("rankings_count", len(rankings)),
		logx.Field("week_start", weekStart.Format("2006-01-02")),
		logx.Field("week_end", weekEnd.Format("2006-01-02")))

	return &doctor.GetWeeklyRankingResp{
		Code:      200,
		Message:   "获取成功",
		Rankings:  rankings,
		WeekStart: weekStart.Format("2006-01-02"),
		WeekEnd:   weekEnd.Format("2006-01-02"),
	}, nil
}

// getWeeklyRankingFromRedis 从Redis获取排行榜数据
// 从Redis有序集合中获取指定数量的排行榜数据，按业绩分数降序排列
func (l *GetWeeklyRankingLogic) getWeeklyRankingFromRedis(limit int) ([]*doctor.RankingInfo, error) {
	redisKey := "weekly_ranking"

	// 从Redis获取排行榜数据（按分数降序）
	result, err := l.svcCtx.RedisClient.ZRevRangeWithScores(context.Background(), redisKey, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, fmt.Errorf("从Redis获取排行榜失败: %v", err)
	}

	var rankings []*doctor.RankingInfo
	for i, item := range result {
		// 解析成员数据：doctor_id:doctor_name:dept_name:prescription_count
		member := item.Member.(string)
		parts := strings.Split(member, ":")
		if len(parts) >= 4 {
			doctorID, _ := strconv.ParseInt(parts[0], 10, 64)
			doctorName := parts[1]
			deptName := parts[2]
			prescriptionCount, _ := strconv.ParseInt(parts[3], 10, 64)

			// 构造排行榜信息
			ranking := &doctor.RankingInfo{
				Rank:              int64(i + 1),
				DoctorId:          doctorID,
				DoctorName:        doctorName,
				DeptName:          deptName,
				TotalPerformance:  item.Score,
				PrescriptionCount: prescriptionCount,
			}
			rankings = append(rankings, ranking)
		}
	}

	return rankings, nil
}

// getWeekStartEnd 获取本周的开始和结束时间
// 返回本周一00:00:00到周日23:59:59的时间范围
func (l *GetWeeklyRankingLogic) getWeekStartEnd() (time.Time, time.Time) {
	now := time.Now()
	weekday := now.Weekday()
	// 调整周日为7，其他日期减1
	if weekday == time.Sunday {
		weekday = 7
	} else {
		weekday--
	}
	// 计算本周一的日期
	weekStart := now.AddDate(0, 0, -int(weekday))
	// 设置为本周一00:00:00
	weekStart = time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, weekStart.Location())
	// 计算本周日的日期
	weekEnd := weekStart.AddDate(0, 0, 6)
	// 设置为本周日23:59:59
	weekEnd = time.Date(weekEnd.Year(), weekEnd.Month(), weekEnd.Day(), 23, 59, 59, 999999999, weekEnd.Location())
	return weekStart, weekEnd
}
