package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePrescriptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePrescriptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePrescriptionLogic {
	return &CreatePrescriptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePrescriptionLogic) CreatePrescription(req *types.CreatePrescriptionReq) (resp *types.CreatePrescriptionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
