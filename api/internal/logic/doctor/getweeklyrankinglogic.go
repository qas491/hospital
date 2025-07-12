package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"
	"github.com/qas491/hospital/doctor_srv/doctor"

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

func (l *GetWeeklyRankingLogic) GetWeeklyRanking(req *types.GetWeeklyRankingReq) (*types.GetWeeklyRankingResp, error) {
	rpcResp, err := l.svcCtx.DoctorRpc.GetWeeklyRanking(l.ctx, &doctor.GetWeeklyRankingReq{
		Limit: req.Limit,
	})
	if err != nil {
		return nil, err
	}

	var rankings []types.RankingInfo
	for _, r := range rpcResp.Rankings {
		rankings = append(rankings, types.RankingInfo{
			Rank:               r.Rank,
			Doctor_id:          r.DoctorId,
			Doctor_name:        r.DoctorName,
			Dept_name:          r.DeptName,
			Total_performance:  r.TotalPerformance,
			Prescription_count: r.PrescriptionCount,
		})
	}

	return &types.GetWeeklyRankingResp{
		Code:       rpcResp.Code,
		Message:    rpcResp.Message,
		Rankings:   rankings,
		Week_start: rpcResp.WeekStart,
		Week_end:   rpcResp.WeekEnd,
	}, nil
}
