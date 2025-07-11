package logic

import (
	"context"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/qas491/hospital/doctor_srv/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateWeeklyRankingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateWeeklyRankingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateWeeklyRankingLogic {
	return &GenerateWeeklyRankingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 生成周排行榜
func (l *GenerateWeeklyRankingLogic) GenerateWeeklyRanking(in *doctor.GenerateWeeklyRankingReq) (*doctor.GenerateWeeklyRankingResp, error) {
	// todo: add your logic here and delete this line

	return &doctor.GenerateWeeklyRankingResp{}, nil
}
