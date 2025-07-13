package logic // 逻辑层包名

import ( // 导入所需的包
	"context" // 上下文包，用于传递请求上下文
	"fmt"     // 格式化输出包
	"strconv" // 字符串与数字转换包
	"strings" // 字符串处理包
	"time"    // 时间处理包

	"github.com/qas491/hospital/api/internal/svc"   // 服务上下文包
	"github.com/qas491/hospital/api/internal/types" // 类型定义包
	"github.com/qas491/hospital/api/model/redis"    // Redis操作包
	"github.com/zeromicro/go-zero/core/logx"        // 日志包
)

// GetWeeklyRankingLogic 用于处理获取周排行榜的业务逻辑
// 结构体包含日志、上下文和服务上下文
// 这是业务逻辑的核心结构体
// 用于封装与周排行榜相关的操作
// 便于后续扩展和维护
type GetWeeklyRankingLogic struct { // 定义结构体
	logx.Logger                     // 日志记录器
	ctx         context.Context     // 请求上下文
	svcCtx      *svc.ServiceContext // 服务上下文
}

// NewGetWeeklyRankingLogic 创建 GetWeeklyRankingLogic 实例
// 参数为请求上下文和服务上下文
// 返回一个新的 GetWeeklyRankingLogic 指针
func NewGetWeeklyRankingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWeeklyRankingLogic {
	return &GetWeeklyRankingLogic{ // 返回结构体实例
		Logger: logx.WithContext(ctx), // 初始化日志
		ctx:    ctx,                   // 赋值上下文
		svcCtx: svcCtx,                // 赋值服务上下文
	}
}

// GetWeeklyRanking 获取排行榜数据
// 参数为请求体，返回响应体和错误信息
func (l *GetWeeklyRankingLogic) GetWeeklyRanking(req *types.GetWeeklyRankingReq) (resp *types.GetWeeklyRankingResp, err error) {
	// 从Redis获取排行榜数据
	rankings, err := l.getWeeklyRankingFromRedis(int(req.Limit)) // 调用内部方法获取排行榜
	if err != nil {                                              // 如果有错误
		return nil, fmt.Errorf("获取排行榜失败: %v", err) // 返回错误信息
	}

	// 构建响应
	resp = &types.GetWeeklyRankingResp{ // 初始化响应结构体
		Code:       200,                                   // 状态码
		Message:    "获取成功",                                // 提示信息
		Rankings:   rankings,                              // 排行榜数据
		Week_start: l.getWeekStart().Format("2006-01-02"), // 本周开始日期
		Week_end:   l.getWeekEnd().Format("2006-01-02"),   // 本周结束日期
	}

	return resp, nil // 返回响应和空错误
}

// getWeeklyRankingFromRedis 从Redis获取排行榜数据
// 参数为排行榜条数，返回排行榜信息数组和错误
func (l *GetWeeklyRankingLogic) getWeeklyRankingFromRedis(limit int) ([]types.RankingInfo, error) {
	redisKey := "weekly_ranking" // Redis中排行榜的key

	// 从Redis获取排行榜数据（按分数降序）
	result, err := redis.RDB.ZRevRangeWithScores(context.Background(), redisKey, 0, int64(limit-1)).Result() // 获取数据
	if err != nil {                                                                                          // 如果有错误
		return nil, fmt.Errorf("从Redis获取排行榜失败: %v", err) // 返回错误
	}

	var rankings []types.RankingInfo // 定义排行榜切片
	for i, item := range result {    // 遍历结果
		// 解析成员数据
		member := item.Member.(string)      // 获取成员字符串
		parts := strings.Split(member, ":") // 按冒号分割
		if len(parts) >= 4 {                // 判断分割后长度
			doctorID, _ := strconv.ParseInt(parts[0], 10, 64)          // 医生ID
			doctorName := parts[1]                                     // 医生姓名
			deptName := parts[2]                                       // 科室名称
			prescriptionCount, _ := strconv.ParseInt(parts[3], 10, 64) // 处方数量

			ranking := types.RankingInfo{ // 构造排行榜项
				Rank:               int64(i + 1),      // 排名
				Doctor_id:          doctorID,          // 医生ID
				Doctor_name:        doctorName,        // 医生姓名
				Dept_name:          deptName,          // 科室名称
				Total_performance:  item.Score,        // 总业绩
				Prescription_count: prescriptionCount, // 处方数量
			}
			rankings = append(rankings, ranking) // 添加到切片
		}
	}

	return rankings, nil // 返回排行榜和空错误
}

// getWeekStart 获取本周开始时间
// 返回本周一的零点时间
func (l *GetWeeklyRankingLogic) getWeekStart() time.Time {
	now := time.Now()           // 当前时间
	weekday := now.Weekday()    // 当前星期几
	if weekday == time.Sunday { // 如果是周日
		weekday = 7 // 设为7
	} else {
		weekday-- // 其它情况减一
	}
	weekStart := now.AddDate(0, 0, -int(weekday))                                                            // 计算本周一日期
	return time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, weekStart.Location()) // 返回本周一零点
}

// getWeekEnd 获取本周结束时间
// 返回本周日的23:59:59时间
func (l *GetWeeklyRankingLogic) getWeekEnd() time.Time {
	weekStart := l.getWeekStart()                                                                               // 获取本周一
	weekEnd := weekStart.AddDate(0, 0, 6)                                                                       // 加6天为本周日
	return time.Date(weekEnd.Year(), weekEnd.Month(), weekEnd.Day(), 23, 59, 59, 999999999, weekEnd.Location()) // 返回本周日23:59:59
}
