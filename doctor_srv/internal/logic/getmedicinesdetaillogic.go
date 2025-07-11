package logic

import (
	"context"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/qas491/hospital/doctor_srv/internal/svc"

	"github.com/qas491/hospital/doctor_srv/model/mysql"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMedicinesDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMedicinesDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMedicinesDetailLogic {
	return &GetMedicinesDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取药品详情
func (l *GetMedicinesDetailLogic) GetMedicinesDetail(in *doctor.GetMedicinesDetailReq) (*doctor.GetMedicinesDetailResp, error) {
	var med mysql.StockMedicines
	err := l.svcCtx.DB.Where("medicines_id = ?", in.MedicinesId).First(&med).Error
	if err != nil {
		return &doctor.GetMedicinesDetailResp{
			Code:    1,
			Message: "未找到药品: " + err.Error(),
		}, nil
	}
	resp := &doctor.GetMedicinesDetailResp{
		Code:    0,
		Message: "ok",
		Detail: &doctor.MedicinesInfo{
			MedicinesId:             med.MedicinesID,
			MedicinesNumber:         med.MedicinesNumber,
			MedicinesName:           med.MedicinesName,
			MedicinesType:           med.MedicinesType,
			PrescriptionType:        med.PrescriptionType,
			PrescriptionPrice:       med.PrescriptionPrice,
			Unit:                    med.Unit,
			Conversion:              int32(med.Conversion),
			Keywords:                med.Keywords,
			ProducterId:             med.ProducterID,
			Status:                  med.Status,
			MedicinesStockNum:       med.MedicinesStockNum,
			MedicinesStockDangerNum: med.MedicinesStockDangerNum,
			CreateTime:              FormatTime(med.CreateTime),
			UpdateTime:              FormatTime(med.UpdateTime),
			CreateBy:                med.CreateBy,
			UpdateBy:                med.UpdateBy,
			DelFlag:                 med.DelFlag,
		},
	}
	return resp, nil
}
