package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/qas491/hospital/doctor_srv/doctor"
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

func (l *GetPrescriptionListLogic) GetPrescriptionList(req *types.GetPrescriptionListReq) (*types.GetPrescriptionListResp, error) {
	rpcResp, err := l.svcCtx.DoctorRpc.GetPrescriptionList(l.ctx, &doctor.GetPrescriptionListReq{
		Page:        req.Page,
		PageSize:    req.Page_size,
		PatientId:   req.Patient_id,
		PatientName: req.Patient_name,
		UserId:      req.User_id,
		CoType:      req.Co_type,
		Status:      req.Status,
		StartTime:   req.Start_time,
		EndTime:     req.End_time,
	})
	if err != nil {
		return nil, err
	}
	var list []types.PrescriptionInfo
	for _, p := range rpcResp.List {
		list = append(list, convertToPrescriptionInfo(p))
	}
	return &types.GetPrescriptionListResp{
		Code:    rpcResp.Code,
		Message: rpcResp.Message,
		List:    list,
		Total:   rpcResp.Total,
	}, nil
}
