package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMedicinesDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMedicinesDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMedicinesDetailLogic {
	return &GetMedicinesDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMedicinesDetailLogic) GetMedicinesDetail(req *types.GetMedicinesDetailReq) (*types.GetMedicinesDetailResp, error) {
	rpcResp, err := l.svcCtx.DoctorRpc.GetMedicinesDetail(l.ctx, &doctor.GetMedicinesDetailReq{
		MedicinesId: req.Medicines_id,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetMedicinesDetailResp{
		Code:    rpcResp.Code,
		Message: rpcResp.Message,
		Detail:  convertToMedicinesInfo(rpcResp.Detail),
	}, nil
}

func convertToMedicinesInfo(in *doctor.MedicinesInfo) types.MedicinesInfo {
	if in == nil {
		return types.MedicinesInfo{}
	}
	return types.MedicinesInfo{
		Medicines_id:               in.MedicinesId,
		Medicines_number:           in.MedicinesNumber,
		Medicines_name:             in.MedicinesName,
		Medicines_type:             in.MedicinesType,
		Prescription_type:          in.PrescriptionType,
		Prescription_price:         in.PrescriptionPrice,
		Unit:                       in.Unit,
		Conversion:                 in.Conversion,
		Keywords:                   in.Keywords,
		Producter_id:               in.ProducterId,
		Status:                     in.Status,
		Medicines_stock_num:        in.MedicinesStockNum,
		Medicines_stock_danger_num: in.MedicinesStockDangerNum,
		Create_time:                in.CreateTime,
		Update_time:                in.UpdateTime,
		Create_by:                  in.CreateBy,
		Update_by:                  in.UpdateBy,
		Del_flag:                   in.DelFlag,
	}
}
