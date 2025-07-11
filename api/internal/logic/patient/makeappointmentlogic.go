package patient

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MakeAppointmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMakeAppointmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakeAppointmentLogic {
	return &MakeAppointmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MakeAppointmentLogic) MakeAppointment(req *types.MakeAppointmentReq) (resp *types.MakeAppointmentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
