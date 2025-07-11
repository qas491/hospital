package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/qas491/hospital/doctor_srv/internal/svc"
	"github.com/qas491/hospital/doctor_srv/model/mysql"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMedicinesListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMedicinesListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMedicinesListLogic {
	return &GetMedicinesListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetMedicinesList 获取药品列表接口
// 支持分页查询和多种条件筛选
func (l *GetMedicinesListLogic) GetMedicinesList(in *doctor.GetMedicinesListReq) (*doctor.GetMedicinesListResp, error) {
	// 1. 参数验证
	if err := l.validateListRequest(in); err != nil {
		return &doctor.GetMedicinesListResp{
			Code:    400,
			Message: err.Error(),
		}, nil
	}

	// 2. 构建查询条件
	query := l.svcCtx.DB.Model(&mysql.StockMedicines{}).Where("del_flag = '0'")

	// 添加查询条件
	if in.MedicinesName != "" {
		query = query.Where("medicines_name LIKE ?", "%"+in.MedicinesName+"%")
	}
	if in.MedicinesType != "" {
		query = query.Where("medicines_type = ?", in.MedicinesType)
	}
	if in.PrescriptionType != "" {
		query = query.Where("prescription_type = ?", in.PrescriptionType)
	}
	if in.Status != "" {
		query = query.Where("status = ?", in.Status)
	}
	if in.Keywords != "" {
		query = query.Where("keywords LIKE ? OR medicines_name LIKE ?",
			"%"+in.Keywords+"%", "%"+in.Keywords+"%")
	}

	// 3. 获取总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return &doctor.GetMedicinesListResp{
			Code:    500,
			Message: "查询药品总数失败",
		}, nil
	}

	// 4. 分页查询
	var medicines []mysql.StockMedicines
	offset := (in.Page - 1) * in.PageSize
	if err := query.Offset(int(offset)).Limit(int(in.PageSize)).
		Order("medicines_id DESC").Find(&medicines).Error; err != nil {
		return &doctor.GetMedicinesListResp{
			Code:    500,
			Message: "查询药品列表失败",
		}, nil
	}

	// 5. 构建响应数据
	var medicinesList []*doctor.MedicinesInfo
	for _, medicine := range medicines {
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
		medicinesList = append(medicinesList, medicinesInfo)
	}

	// 6. 记录查询日志
	go l.recordQueryLog(in.Page, in.PageSize, len(medicinesList))

	return &doctor.GetMedicinesListResp{
		Code:    200,
		Message: "获取药品列表成功",
		List:    medicinesList,
		Total:   total,
	}, nil
}

// validateListRequest 验证列表请求参数
func (l *GetMedicinesListLogic) validateListRequest(in *doctor.GetMedicinesListReq) error {
	if in.Page <= 0 {
		return fmt.Errorf("页码必须大于0")
	}
	if in.PageSize <= 0 || in.PageSize > 100 {
		return fmt.Errorf("每页数量必须在1-100之间")
	}
	return nil
}

// formatTime 格式化时间
func (l *GetMedicinesListLogic) formatTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

// recordQueryLog 记录查询日志
func (l *GetMedicinesListLogic) recordQueryLog(page, pageSize int64, resultCount int) {
	log := &mysql.SysOperLog{
		Title:        "查询药品列表",
		BusinessType: "QUERY_MEDICINES_LIST",
		Method:       "GetMedicinesList",
		OperName:     "doctor", // 可以从上下文获取当前用户
		OperURL:      "/doctor/medicines/list",
		OperParam:    fmt.Sprintf("页码: %d, 每页数量: %d, 返回数量: %d", page, pageSize, resultCount),
		Status:       "0", // 成功
		OperTime:     &time.Time{},
	}

	if err := l.svcCtx.DB.Create(log).Error; err != nil {
		l.Logger.Errorf("记录查询日志失败: %v", err)
	}
}
