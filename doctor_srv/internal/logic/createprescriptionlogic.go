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

type CreatePrescriptionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePrescriptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePrescriptionLogic {
	return &CreatePrescriptionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreatePrescription 开处方接口
// 实现高并发库存管理，确保库存不会超卖
func (l *CreatePrescriptionLogic) CreatePrescription(in *doctor.CreatePrescriptionReq) (*doctor.CreatePrescriptionResp, error) {
	// 1. 参数验证
	if err := l.validatePrescriptionRequest(in); err != nil {
		return &doctor.CreatePrescriptionResp{
			Code:    400,
			Message: err.Error(),
		}, nil
	}

	// 2. 使用事务确保数据一致性
	tx := l.svcCtx.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 3. 检查并锁定库存（使用悲观锁防止超卖）
	if err := l.checkAndLockInventory(tx, in.Items); err != nil {
		tx.Rollback()
		return &doctor.CreatePrescriptionResp{
			Code:    400,
			Message: fmt.Sprintf("库存不足: %v", err),
		}, nil
	}

	// 4. 创建处方主表记录
	careOrder := &mysql.HisCareOrder{
		CoID:        in.CoId,
		CoType:      in.CoType,
		UserID:      in.UserId,
		PatientID:   in.PatientId,
		PatientName: in.PatientName,
		ChID:        in.ChId,
		AllAmount:   in.AllAmount,
		CreateBy:    in.CreateBy,
		CreateTime:  &time.Time{},
	}

	if err := tx.Create(careOrder).Error; err != nil {
		tx.Rollback()
		return &doctor.CreatePrescriptionResp{
			Code:    500,
			Message: "创建处方失败",
		}, nil
	}

	// 5. 创建处方明细记录
	for _, item := range in.Items {
		orderItem := &mysql.HisCareOrderItem{
			ItemID:     item.ItemId,
			CoID:       in.CoId,
			ItemRefID:  item.ItemRefId,
			ItemName:   item.ItemName,
			ItemType:   item.ItemType,
			Num:        item.Num,
			Price:      item.Price,
			Amount:     item.Amount,
			Remark:     item.Remark,
			Status:     item.Status,
			CreateTime: &time.Time{},
		}

		if err := tx.Create(orderItem).Error; err != nil {
			tx.Rollback()
			return &doctor.CreatePrescriptionResp{
				Code:    500,
				Message: "创建处方明细失败",
			}, nil
		}

		// 6. 扣减库存（在事务中执行）
		if err := l.deductInventory(tx, item.ItemRefId, item.Num); err != nil {
			tx.Rollback()
			return &doctor.CreatePrescriptionResp{
				Code:    500,
				Message: "扣减库存失败",
			}, nil
		}
	}

	// 7. 提交事务
	if err := tx.Commit().Error; err != nil {
		return &doctor.CreatePrescriptionResp{
			Code:    500,
			Message: "提交事务失败",
		}, nil
	}

	// 8. 记录操作日志
	go l.recordOperationLog("CREATE_PRESCRIPTION", in.CoId, in.CreateBy)

	return &doctor.CreatePrescriptionResp{
		Code:    200,
		Message: "开处方成功",
		CoId:    in.CoId,
	}, nil
}

// validatePrescriptionRequest 验证处方请求参数
func (l *CreatePrescriptionLogic) validatePrescriptionRequest(in *doctor.CreatePrescriptionReq) error {
	if in.CoId == "" {
		return fmt.Errorf("处方ID不能为空")
	}
	if in.PatientId == "" {
		return fmt.Errorf("患者ID不能为空")
	}
	if in.UserId == 0 {
		return fmt.Errorf("医生ID不能为空")
	}
	if len(in.Items) == 0 {
		return fmt.Errorf("处方项目不能为空")
	}
	return nil
}

// checkAndLockInventory 检查并锁定库存
func (l *CreatePrescriptionLogic) checkAndLockInventory(tx *gorm.DB, items []*doctor.PrescriptionItem) error {
	for _, item := range items {
		// 使用SELECT FOR UPDATE锁定库存记录
		var medicine mysql.StockMedicines
		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Where("medicines_id = ? AND del_flag = '0'", item.ItemRefId).
			First(&medicine).Error; err != nil {
			return fmt.Errorf("药品不存在: %s", item.ItemRefId)
		}

		// 检查库存是否充足
		if medicine.MedicinesStockNum < item.Num {
			return fmt.Errorf("药品 %s 库存不足", medicine.MedicinesName)
		}
	}
	return nil
}

// deductInventory 扣减库存
func (l *CreatePrescriptionLogic) deductInventory(tx *gorm.DB, medicinesId string, num float64) error {
	// 使用原子操作扣减库存
	result := tx.Model(&mysql.StockMedicines{}).
		Where("medicines_id = ? AND medicines_stock_num >= ? AND del_flag = '0'", medicinesId, num).
		Update("medicines_stock_num", gorm.Expr("medicines_stock_num - ?", num))

	if result.Error != nil {
		return fmt.Errorf("扣减库存失败: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("库存不足或药品不存在")
	}

	return nil
}

// recordOperationLog 记录操作日志
func (l *CreatePrescriptionLogic) recordOperationLog(operation, targetId, operator string) {
	log := &mysql.SysOperLog{
		Title:        "开处方",
		BusinessType: operation,
		Method:       "CreatePrescription",
		OperName:     operator,
		OperURL:      "/doctor/prescription/create",
		OperParam:    fmt.Sprintf("处方ID: %s", targetId),
		Status:       "0", // 成功
		OperTime:     &time.Time{},
	}

	if err := l.svcCtx.DB.Create(log).Error; err != nil {
		l.Logger.Errorf("记录操作日志失败: %v", err)
	}
}
