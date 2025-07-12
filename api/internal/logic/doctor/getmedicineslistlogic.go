package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMedicinesListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMedicinesListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMedicinesListLogic {
	return &GetMedicinesListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMedicinesListLogic) GetMedicinesList(req *types.GetMedicinesListReq) (*types.GetMedicinesListResp, error) {
	rpcResp, err := l.svcCtx.DoctorRpc.GetMedicinesList(l.ctx, &doctor.GetMedicinesListReq{
		Page:             req.Page,
		PageSize:         req.Page_size,
		MedicinesName:    req.Medicines_name,
		MedicinesType:    req.Medicines_type,
		PrescriptionType: req.Prescription_type,
		Status:           req.Status,
		Keywords:         req.Keywords,
	})
	if err != nil {
		return nil, err
	}
	var list []types.MedicinesInfo
	for _, m := range rpcResp.List {
		list = append(list, convertToMedicinesInfo(m))
	}
	return &types.GetMedicinesListResp{
		Code:    rpcResp.Code,
		Message: rpcResp.Message,
		List:    list,
		Total:   rpcResp.Total,
	}, nil
}
