package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"
	"github.com/qas491/hospital/doctor_srv/doctor"

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

func (l *CreatePrescriptionLogic) CreatePrescription(req *types.CreatePrescriptionReq) (*types.CreatePrescriptionResp, error) {
	rpcReq := &doctor.CreatePrescriptionReq{
		CoId:        req.Co_id,
		CoType:      req.Co_type,
		UserId:      req.User_id,
		PatientId:   req.Patient_id,
		PatientName: req.Patient_name,
		ChId:        req.Ch_id,
		AllAmount:   req.All_amount,
		CreateBy:    req.Create_by,
		Items:       convertToDoctorPrescriptionItems(req.Items),
	}
	rpcResp, err := l.svcCtx.DoctorRpc.CreatePrescription(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	return &types.CreatePrescriptionResp{
		Code:    rpcResp.Code,
		Message: rpcResp.Message,
		Co_id:   rpcResp.CoId,
	}, nil
}

func convertToDoctorPrescriptionItems(items []types.PrescriptionItem) []*doctor.PrescriptionItem {
	var result []*doctor.PrescriptionItem
	for _, item := range items {
		result = append(result, &doctor.PrescriptionItem{
			ItemId:    item.Item_id,
			ItemRefId: item.Item_ref_id,
			ItemName:  item.Item_name,
			ItemType:  item.Item_type,
			Num:       item.Num,
			Price:     item.Price,
			Amount:    item.Amount,
			Remark:    item.Remark,
			Status:    item.Status,
		})
	}
	return result
}
