package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReviewPrescriptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReviewPrescriptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReviewPrescriptionLogic {
	return &ReviewPrescriptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReviewPrescriptionLogic) ReviewPrescription(req *types.ReviewPrescriptionReq) (resp *types.ReviewPrescriptionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
