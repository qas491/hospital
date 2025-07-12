package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/qas491/hospital/doctor_srv/doctor"
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

func (l *GetPrescriptionDetailLogic) GetPrescriptionDetail(req *types.GetPrescriptionDetailReq) (*types.GetPrescriptionDetailResp, error) {
	rpcResp, err := l.svcCtx.DoctorRpc.GetPrescriptionDetail(l.ctx, &doctor.GetPrescriptionDetailReq{
		CoId: req.Co_id,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetPrescriptionDetailResp{
		Code:    rpcResp.Code,
		Message: rpcResp.Message,
		Detail:  convertToPrescriptionInfo(rpcResp.Detail),
		Qrcode:  rpcResp.Qrcode,
	}, nil
}

// 类型转换辅助函数
func convertToPrescriptionInfo(in *doctor.PrescriptionInfo) types.PrescriptionInfo {
	if in == nil {
		return types.PrescriptionInfo{}
	}
	var items []types.PrescriptionItem
	for _, item := range in.Items {
		items = append(items, types.PrescriptionItem{
			Item_id:     item.ItemId,
			Item_ref_id: item.ItemRefId,
			Item_name:   item.ItemName,
			Item_type:   item.ItemType,
			Num:         item.Num,
			Price:       item.Price,
			Amount:      item.Amount,
			Remark:      item.Remark,
			Status:      item.Status,
		})
	}
	return types.PrescriptionInfo{
		Co_id:        in.CoId,
		Co_type:      in.CoType,
		User_id:      in.UserId,
		Patient_id:   in.PatientId,
		Patient_name: in.PatientName,
		Ch_id:        in.ChId,
		All_amount:   in.AllAmount,
		Create_by:    in.CreateBy,
		Create_time:  in.CreateTime,
		Update_by:    in.UpdateBy,
		Update_time:  in.UpdateTime,
		Items:        items,
	}
}
