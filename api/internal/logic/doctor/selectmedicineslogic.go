package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/zeromicro/go-zero/core/logx"
)

type SelectMedicinesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectMedicinesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectMedicinesLogic {
	return &SelectMedicinesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectMedicinesLogic) SelectMedicines(req *types.SelectMedicinesReq) (*types.SelectMedicinesResp, error) {
	rpcResp, err := l.svcCtx.DoctorRpc.SelectMedicines(l.ctx, &doctor.SelectMedicinesReq{
		MedicinesId:      req.Medicines_id,
		MedicinesName:    req.Medicines_name,
		MedicinesType:    req.Medicines_type,
		PrescriptionType: req.Prescription_type,
		Num:              req.Num,
		Unit:             req.Unit,
	})
	if err != nil {
		return nil, err
	}
	return &types.SelectMedicinesResp{
		Code:      rpcResp.Code,
		Message:   rpcResp.Message,
		Medicines: convertToMedicinesInfo(rpcResp.Medicines),
	}, nil
}
