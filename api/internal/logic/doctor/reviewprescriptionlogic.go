package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"
	"github.com/qas491/hospital/doctor_srv/doctor"

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

func (l *ReviewPrescriptionLogic) ReviewPrescription(req *types.ReviewPrescriptionReq) (*types.ReviewPrescriptionResp, error) {
	rpcResp, err := l.svcCtx.DoctorRpc.ReviewPrescription(l.ctx, &doctor.ReviewPrescriptionReq{
		CoId:         req.Co_id,
		ReviewStatus: req.Review_status,
		ReviewBy:     req.Review_by,
		ReviewRemark: req.Review_remark,
	})
	if err != nil {
		return nil, err
	}
	return &types.ReviewPrescriptionResp{
		Code:    rpcResp.Code,
		Message: rpcResp.Message,
		Success: rpcResp.Success,
	}, nil
}
