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

type ReviewPrescriptionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReviewPrescriptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReviewPrescriptionLogic {
	return &ReviewPrescriptionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ReviewPrescription 审核处方接口
func (l *ReviewPrescriptionLogic) ReviewPrescription(in *doctor.ReviewPrescriptionReq) (*doctor.ReviewPrescriptionResp, error) {
	// 1. 参数验证
	if err := l.validateReviewRequest(in); err != nil {
		return &doctor.ReviewPrescriptionResp{
			Code:    400,
			Message: err.Error(),
		}, nil
	}

	// 2. 检查处方是否存在
	var careOrder mysql.HisCareOrder
	if err := l.svcCtx.DB.Where("co_id = ?", in.CoId).First(&careOrder).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &doctor.ReviewPrescriptionResp{
				Code:    404,
				Message: "处方不存在",
			}, nil
		}
		return &doctor.ReviewPrescriptionResp{
			Code:    500,
			Message: "查询处方失败",
		}, nil
	}

	// 3. 使用事务进行审核操作
	tx := l.svcCtx.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 4. 更新处方状态
	updateData := map[string]interface{}{
		"update_by":   in.ReviewBy,
		"update_time": &time.Time{},
	}

	// 根据审核状态设置相应的字段
	switch in.ReviewStatus {
	case "APPROVED":
		// 由于HisCareOrder没有Status字段，我们可以在备注中记录状态
		updateData["remark"] = "APPROVED"
	case "REJECTED":
		updateData["remark"] = "REJECTED: " + in.ReviewRemark
	case "PENDING":
		updateData["remark"] = "PENDING"
	default:
		return &doctor.ReviewPrescriptionResp{
			Code:    400,
			Message: "无效的审核状态",
		}, nil
	}

	if err := tx.Model(&mysql.HisCareOrder{}).
		Where("co_id = ?", in.CoId).
		Updates(updateData).Error; err != nil {
		tx.Rollback()
		return &doctor.ReviewPrescriptionResp{
			Code:    500,
			Message: "更新处方状态失败",
		}, nil
	}

	// 5. 如果审核通过，更新处方明细状态
	if in.ReviewStatus == "APPROVED" {
		if err := l.updatePrescriptionItemsStatus(tx, in.CoId, "APPROVED"); err != nil {
			tx.Rollback()
			return &doctor.ReviewPrescriptionResp{
				Code:    500,
				Message: "更新处方明细状态失败",
			}, nil
		}
	}

	// 6. 如果审核拒绝，恢复库存
	if in.ReviewStatus == "REJECTED" {
		if err := l.restoreInventory(tx, in.CoId); err != nil {
			tx.Rollback()
			return &doctor.ReviewPrescriptionResp{
				Code:    500,
				Message: "恢复库存失败",
			}, nil
		}
	}

	// 7. 提交事务
	if err := tx.Commit().Error; err != nil {
		return &doctor.ReviewPrescriptionResp{
			Code:    500,
			Message: "提交事务失败",
		}, nil
	}

	// 8. 记录操作日志
	go l.recordReviewLog(in.CoId, in.ReviewStatus, in.ReviewBy, in.ReviewRemark)

	// 9. 发送审核结果通知（异步）
	go l.sendReviewNotification(in.CoId, in.ReviewStatus, careOrder.PatientID)

	return &doctor.ReviewPrescriptionResp{
		Code:    200,
		Message: "审核处方成功",
		Success: true,
	}, nil
}

// validateReviewRequest 验证审核请求参数
func (l *ReviewPrescriptionLogic) validateReviewRequest(in *doctor.ReviewPrescriptionReq) error {
	if in.CoId == "" {
		return fmt.Errorf("处方ID不能为空")
	}
	if in.ReviewStatus == "" {
		return fmt.Errorf("审核状态不能为空")
	}
	if in.ReviewBy == "" {
		return fmt.Errorf("审核人不能为空")
	}
	return nil
}

// updatePrescriptionItemsStatus 更新处方明细状态
func (l *ReviewPrescriptionLogic) updatePrescriptionItemsStatus(tx *gorm.DB, coId, status string) error {
	return tx.Model(&mysql.HisCareOrderItem{}).
		Where("co_id = ?", coId).
		Update("status", status).Error
}

// restoreInventory 恢复库存（审核拒绝时）
func (l *ReviewPrescriptionLogic) restoreInventory(tx *gorm.DB, coId string) error {
	// 查询处方明细
	var items []mysql.HisCareOrderItem
	if err := tx.Where("co_id = ?", coId).Find(&items).Error; err != nil {
		return err
	}

	// 恢复每个药品的库存
	for _, item := range items {
		result := tx.Model(&mysql.StockMedicines{}).
			Where("medicines_id = ? AND del_flag = '0'", item.ItemRefID).
			Update("medicines_stock_num", gorm.Expr("medicines_stock_num + ?", item.Num))

		if result.Error != nil {
			return fmt.Errorf("恢复药品 %s 库存失败: %v", item.ItemName, result.Error)
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("药品 %s 不存在或已删除", item.ItemName)
		}
	}

	return nil
}

// recordReviewLog 记录审核日志
func (l *ReviewPrescriptionLogic) recordReviewLog(coId, status, reviewer, remark string) {
	log := &mysql.SysOperLog{
		Title:        "审核处方",
		BusinessType: "REVIEW_PRESCRIPTION",
		Method:       "ReviewPrescription",
		OperName:     reviewer,
		OperURL:      "/doctor/prescription/review",
		OperParam:    fmt.Sprintf("处方ID: %s, 审核状态: %s, 备注: %s", coId, status, remark),
		Status:       "0", // 成功
		OperTime:     &time.Time{},
	}

	if err := l.svcCtx.DB.Create(log).Error; err != nil {
		l.Logger.Errorf("记录审核日志失败: %v", err)
	}
}

// sendReviewNotification 发送审核结果通知
func (l *ReviewPrescriptionLogic) sendReviewNotification(coId, status, patientId string) {
	// 这里可以实现发送短信、邮件或推送通知的逻辑
	// 例如：发送短信通知患者审核结果
	notificationType := "SMS"

	switch status {
	case "APPROVED":
		fmt.Printf("您的处方 %s 已审核通过，请及时取药。\n", coId)
	case "REJECTED":
		fmt.Printf("您的处方 %s 审核未通过，请重新开具处方。\n", coId)
	case "PENDING":
		fmt.Printf("您的处方 %s 正在审核中，请耐心等待。\n", coId)
	}

	// 记录短信发送日志
	smsLog := &mysql.SysSmsLog{
		ID:         time.Now().UnixNano(),
		Mobile:     "", // 需要从患者信息中获取手机号
		CreateTime: &time.Time{},
		Code:       status,
		Status:     "SENT",
		Type:       notificationType,
	}

	if err := l.svcCtx.DB.Create(smsLog).Error; err != nil {
		l.Logger.Errorf("记录短信发送日志失败: %v", err)
	}

	l.Logger.Infof("发送审核通知: 患者ID=%s, 处方ID=%s, 状态=%s", patientId, coId, status)
}
