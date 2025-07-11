package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrescriptionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPrescriptionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrescriptionListLogic {
	return &GetPrescriptionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPrescriptionListLogic) GetPrescriptionList(req *types.GetPrescriptionListReq) (resp *types.GetPrescriptionListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
