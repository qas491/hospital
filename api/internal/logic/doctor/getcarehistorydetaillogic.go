package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCareHistoryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCareHistoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCareHistoryDetailLogic {
	return &GetCareHistoryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCareHistoryDetailLogic) GetCareHistoryDetail(req *types.GetCareHistoryDetailReq) (resp *types.GetCareHistoryDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
