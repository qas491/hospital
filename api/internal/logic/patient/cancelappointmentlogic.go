package patient

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelAppointmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelAppointmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelAppointmentLogic {
	return &CancelAppointmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelAppointmentLogic) CancelAppointment() (resp *types.CancelAppointmentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
