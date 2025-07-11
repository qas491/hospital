package logic

import (
	"context"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/qas491/hospital/doctor_srv/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWeeklyRankingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWeeklyRankingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWeeklyRankingLogic {
	return &GetWeeklyRankingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取周排行榜
func (l *GetWeeklyRankingLogic) GetWeeklyRanking(in *doctor.GetWeeklyRankingReq) (*doctor.GetWeeklyRankingResp, error) {
	// todo: add your logic here and delete this line

	return &doctor.GetWeeklyRankingResp{}, nil
}
