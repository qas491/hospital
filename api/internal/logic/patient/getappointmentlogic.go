package patient

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppointmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAppointmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppointmentLogic {
	return &GetAppointmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAppointmentLogic) GetAppointment() (resp *types.GetAppointmentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
