package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrescriptionDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPrescriptionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrescriptionDetailLogic {
	return &GetPrescriptionDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPrescriptionDetailLogic) GetPrescriptionDetail(req *types.GetPrescriptionDetailReq) (resp *types.GetPrescriptionDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
