package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/qas491/hospital/doctor_srv/internal/svc"
	"github.com/qas491/hospital/doctor_srv/model/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type SelectMedicinesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectMedicinesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectMedicinesLogic {
	return &SelectMedicinesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SelectMedicines 选择药品接口
// 支持按药品ID、名称、类型等条件选择药品，并检查库存状态
func (l *SelectMedicinesLogic) SelectMedicines(in *doctor.SelectMedicinesReq) (*doctor.SelectMedicinesResp, error) {
	// 1. 参数验证
	if err := l.validateSelectRequest(in); err != nil {
		return &doctor.SelectMedicinesResp{
			Code:    400,
			Message: err.Error(),
		}, nil
	}

	// 2. 查询药品信息
	var medicine mysql.StockMedicines
	query := l.svcCtx.DB.Model(&mysql.StockMedicines{}).Where("del_flag = '0'")

	// 根据不同的查询条件构建查询
	if in.MedicinesId != "" {
		query = query.Where("medicines_id = ?", in.MedicinesId)
	} else if in.MedicinesName != "" {
		query = query.Where("medicines_name LIKE ?", "%"+in.MedicinesName+"%")
	} else if in.MedicinesType != "" {
		query = query.Where("medicines_type = ?", in.MedicinesType)
	} else if in.PrescriptionType != "" {
		query = query.Where("prescription_type = ?", in.PrescriptionType)
	} else {
		return &doctor.SelectMedicinesResp{
			Code:    400,
			Message: "请提供药品查询条件",
		}, nil
	}

	if err := query.First(&medicine).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &doctor.SelectMedicinesResp{
				Code:    404,
				Message: "药品不存在",
			}, nil
		}
		return &doctor.SelectMedicinesResp{
			Code:    500,
			Message: "查询药品失败",
		}, nil
	}

	// 3. 检查药品状态
	if err := l.checkMedicinesStatus(medicine); err != nil {
		return &doctor.SelectMedicinesResp{
			Code:    400,
			Message: err.Error(),
		}, nil
	}

	// 4. 检查库存是否充足
	if in.Num > 0 {
		if err := l.checkInventory(medicine, in.Num); err != nil {
			return &doctor.SelectMedicinesResp{
				Code:    400,
				Message: err.Error(),
			}, nil
		}
	}

	// 5. 构建响应数据
	medicinesInfo := &doctor.MedicinesInfo{
		MedicinesId:             uint64(medicine.MedicinesID),
		MedicinesNumber:         medicine.MedicinesNumber,
		MedicinesName:           medicine.MedicinesName,
		MedicinesType:           medicine.MedicinesType,
		PrescriptionType:        medicine.PrescriptionType,
		PrescriptionPrice:       medicine.PrescriptionPrice,
		Unit:                    medicine.Unit,
		Conversion:              int32(medicine.Conversion),
		Keywords:                medicine.Keywords,
		ProducterId:             medicine.ProducterID,
		Status:                  medicine.Status,
		MedicinesStockNum:       medicine.MedicinesStockNum,
		MedicinesStockDangerNum: medicine.MedicinesStockDangerNum,
		CreateTime:              l.formatTime(medicine.CreateTime),
		UpdateTime:              l.formatTime(medicine.UpdateTime),
		CreateBy:                medicine.CreateBy,
		UpdateBy:                medicine.UpdateBy,
		DelFlag:                 medicine.DelFlag,
	}

	// 6. 记录选择日志
	go l.recordSelectionLog(medicine.MedicinesID, in.Num, in.Unit)

	return &doctor.SelectMedicinesResp{
		Code:      200,
		Message:   "选择药品成功",
		Medicines: medicinesInfo,
	}, nil
}

// validateSelectRequest 验证选择药品请求参数
func (l *SelectMedicinesLogic) validateSelectRequest(in *doctor.SelectMedicinesReq) error {
	// 至少需要提供一个查询条件
	if in.MedicinesId == "" && in.MedicinesName == "" &&
		in.MedicinesType == "" && in.PrescriptionType == "" {
		return fmt.Errorf("请提供至少一个查询条件")
	}

	// 如果提供了数量，检查数量是否合理
	if in.Num > 0 {
		if in.Num > 1000 {
			return fmt.Errorf("单次选择药品数量不能超过1000")
		}
	}

	return nil
}

// checkMedicinesStatus 检查药品状态
func (l *SelectMedicinesLogic) checkMedicinesStatus(medicine mysql.StockMedicines) error {
	// 检查药品是否已停用
	if medicine.Status != "1" {
		return fmt.Errorf("药品 %s 已停用", medicine.MedicinesName)
	}

	// 检查药品是否已删除
	if medicine.DelFlag != "0" {
		return fmt.Errorf("药品 %s 已删除", medicine.MedicinesName)
	}

	return nil
}

// checkInventory 检查库存是否充足
func (l *SelectMedicinesLogic) checkInventory(medicine mysql.StockMedicines, requiredNum float64) error {
	// 检查库存是否充足
	if medicine.MedicinesStockNum < requiredNum {
		return fmt.Errorf("药品 %s 库存不足，当前库存: %.2f %s，需要: %.2f %s",
			medicine.MedicinesName, medicine.MedicinesStockNum, medicine.Unit, requiredNum, medicine.Unit)
	}

	// 检查是否达到危险库存
	if medicine.MedicinesStockNum-requiredNum <= medicine.MedicinesStockDangerNum {
		logx.Infof("药品 %s 库存即将不足，当前库存: %.2f，危险库存: %.2f",
			medicine.MedicinesName, medicine.MedicinesStockNum, medicine.MedicinesStockDangerNum)
	}

	return nil
}

// formatTime 格式化时间
func (l *SelectMedicinesLogic) formatTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

// recordSelectionLog 记录选择药品日志
func (l *SelectMedicinesLogic) recordSelectionLog(medicinesId uint64, num float64, unit string) {
	log := &mysql.SysOperLog{
		Title:        "选择药品",
		BusinessType: "SELECT_MEDICINES",
		Method:       "SelectMedicines",
		OperName:     "doctor", // 可以从上下文获取当前用户
		OperURL:      "/doctor/medicines/select",
		OperParam:    fmt.Sprintf("药品ID: %d, 数量: %.2f %s", medicinesId, num, unit),
		Status:       "0", // 成功
		OperTime:     &time.Time{},
	}

	if err := l.svcCtx.DB.Create(log).Error; err != nil {
		l.Logger.Errorf("记录选择药品日志失败: %v", err)
	}
}
