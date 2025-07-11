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

type GetPrescriptionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrescriptionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrescriptionListLogic {
	return &GetPrescriptionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetPrescriptionList 获取处方列表接口
// 支持分页查询和多种条件筛选
func (l *GetPrescriptionListLogic) GetPrescriptionList(in *doctor.GetPrescriptionListReq) (*doctor.GetPrescriptionListResp, error) {
	// 1. 参数验证
	if err := l.validateListRequest(in); err != nil {
		return &doctor.GetPrescriptionListResp{
			Code:    400,
			Message: err.Error(),
		}, nil
	}

	// 2. 构建查询条件
	query := l.svcCtx.DB.Model(&mysql.HisCareOrder{})

	// 添加查询条件
	if in.PatientId != "" {
		query = query.Where("patient_id = ?", in.PatientId)
	}
	if in.PatientName != "" {
		query = query.Where("patient_name LIKE ?", "%"+in.PatientName+"%")
	}
	if in.UserId != 0 {
		query = query.Where("user_id = ?", in.UserId)
	}
	if in.CoType != "" {
		query = query.Where("co_type = ?", in.CoType)
	}
	if in.Status != "" {
		query = query.Where("status = ?", in.Status)
	}
	if in.StartTime != "" {
		query = query.Where("create_time >= ?", in.StartTime)
	}
	if in.EndTime != "" {
		query = query.Where("create_time <= ?", in.EndTime)
	}

	// 3. 获取总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return &doctor.GetPrescriptionListResp{
			Code:    500,
			Message: "查询处方总数失败",
		}, nil
	}

	// 4. 分页查询
	var careOrders []mysql.HisCareOrder
	offset := (in.Page - 1) * in.PageSize
	if err := query.Offset(int(offset)).Limit(int(in.PageSize)).
		Order("create_time DESC").Find(&careOrders).Error; err != nil {
		return &doctor.GetPrescriptionListResp{
			Code:    500,
			Message: "查询处方列表失败",
		}, nil
	}

	// 5. 构建响应数据
	var prescriptionList []*doctor.PrescriptionInfo
	for _, careOrder := range careOrders {
		// 查询处方明细
		var items []mysql.HisCareOrderItem
		if err := l.svcCtx.DB.Where("co_id = ?", careOrder.CoID).Find(&items).Error; err != nil {
			l.Logger.Errorf("查询处方明细失败: %v", err)
			continue
		}

		// 构建处方项目列表
		var prescriptionItems []*doctor.PrescriptionItem
		for _, item := range items {
			prescriptionItem := &doctor.PrescriptionItem{
				ItemId:    item.ItemID,
				ItemRefId: item.ItemRefID,
				ItemName:  item.ItemName,
				ItemType:  item.ItemType,
				Num:       item.Num,
				Price:     item.Price,
				Amount:    item.Amount,
				Remark:    item.Remark,
				Status:    item.Status,
			}
			prescriptionItems = append(prescriptionItems, prescriptionItem)
		}

		prescriptionInfo := &doctor.PrescriptionInfo{
			CoId:        careOrder.CoID,
			CoType:      careOrder.CoType,
			UserId:      careOrder.UserID,
			PatientId:   careOrder.PatientID,
			PatientName: careOrder.PatientName,
			ChId:        careOrder.ChID,
			AllAmount:   careOrder.AllAmount,
			CreateBy:    careOrder.CreateBy,
			CreateTime:  l.formatTime(careOrder.CreateTime),
			UpdateBy:    careOrder.UpdateBy,
			UpdateTime:  l.formatTime(careOrder.UpdateTime),
			Items:       prescriptionItems,
		}
		prescriptionList = append(prescriptionList, prescriptionInfo)
	}

	// 6. 记录查询日志
	go l.recordQueryLog(in.Page, in.PageSize, len(prescriptionList))

	return &doctor.GetPrescriptionListResp{
		Code:    200,
		Message: "获取处方列表成功",
		List:    prescriptionList,
		Total:   total,
	}, nil
}

// validateListRequest 验证列表请求参数
func (l *GetPrescriptionListLogic) validateListRequest(in *doctor.GetPrescriptionListReq) error {
	if in.Page <= 0 {
		return fmt.Errorf("页码必须大于0")
	}
	if in.PageSize <= 0 || in.PageSize > 100 {
		return fmt.Errorf("每页数量必须在1-100之间")
	}
	return nil
}

// formatTime 格式化时间
func (l *GetPrescriptionListLogic) formatTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

// recordQueryLog 记录查询日志
func (l *GetPrescriptionListLogic) recordQueryLog(page, pageSize int64, resultCount int) {
	log := &mysql.SysOperLog{
		Title:        "查询处方列表",
		BusinessType: "QUERY_PRESCRIPTION_LIST",
		Method:       "GetPrescriptionList",
		OperName:     "doctor", // 可以从上下文获取当前用户
		OperURL:      "/doctor/prescription/list",
		OperParam:    fmt.Sprintf("页码: %d, 每页数量: %d, 返回数量: %d", page, pageSize, resultCount),
		Status:       "0", // 成功
		OperTime:     &time.Time{},
	}

	if err := l.svcCtx.DB.Create(log).Error; err != nil {
		l.Logger.Errorf("记录查询日志失败: %v", err)
	}
}
