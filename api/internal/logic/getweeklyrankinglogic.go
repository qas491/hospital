package logic

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"
	"github.com/qas491/hospital/api/model/redis"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetWeeklyRankingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWeeklyRankingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWeeklyRankingLogic {
	return &GetWeeklyRankingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWeeklyRankingLogic) GetWeeklyRanking(req *types.GetWeeklyRankingReq) (resp *types.GetWeeklyRankingResp, err error) {
	// 从Redis获取排行榜数据
	rankings, err := l.getWeeklyRankingFromRedis(int(req.Limit))
	if err != nil {
		return nil, fmt.Errorf("获取排行榜失败: %v", err)
	}

	// 构建响应
	resp = &types.GetWeeklyRankingResp{
		Code:       200,
		Message:    "获取成功",
		Rankings:   rankings,
		Week_start: l.getWeekStart().Format("2006-01-02"),
		Week_end:   l.getWeekEnd().Format("2006-01-02"),
	}

	return resp, nil
}

// getWeeklyRankingFromRedis 从Redis获取排行榜数据
func (l *GetWeeklyRankingLogic) getWeeklyRankingFromRedis(limit int) ([]types.RankingInfo, error) {
	redisKey := "weekly_ranking"

	// 从Redis获取排行榜数据（按分数降序）
	result, err := redis.RDB.ZRevRangeWithScores(context.Background(), redisKey, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, fmt.Errorf("从Redis获取排行榜失败: %v", err)
	}

	var rankings []types.RankingInfo
	for i, item := range result {
		// 解析成员数据
		member := item.Member.(string)
		parts := strings.Split(member, ":")
		if len(parts) >= 4 {
			doctorID, _ := strconv.ParseInt(parts[0], 10, 64)
			doctorName := parts[1]
			deptName := parts[2]
			prescriptionCount, _ := strconv.ParseInt(parts[3], 10, 64)

			ranking := types.RankingInfo{
				Rank:               int64(i + 1),
				Doctor_id:          doctorID,
				Doctor_name:        doctorName,
				Dept_name:          deptName,
				Total_performance:  item.Score,
				Prescription_count: prescriptionCount,
			}
			rankings = append(rankings, ranking)
		}
	}

	return rankings, nil
}

// getWeekStart 获取本周开始时间
func (l *GetWeeklyRankingLogic) getWeekStart() time.Time {
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
func (l *GetWeeklyRankingLogic) getWeekEnd() time.Time {
	weekStart := l.getWeekStart()
	weekEnd := weekStart.AddDate(0, 0, 6)
	return time.Date(weekEnd.Year(), weekEnd.Month(), weekEnd.Day(), 23, 59, 59, 999999999, weekEnd.Location())
}
