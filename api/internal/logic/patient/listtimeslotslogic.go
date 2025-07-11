package patient

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTimeSlotsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTimeSlotsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTimeSlotsLogic {
	return &ListTimeSlotsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTimeSlotsLogic) ListTimeSlots(req *types.ListTimeSlotsReq) (resp *types.ListTimeSlotsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
