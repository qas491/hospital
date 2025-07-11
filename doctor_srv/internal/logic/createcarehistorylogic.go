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

type CreateCareHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCareHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCareHistoryLogic {
	return &CreateCareHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateCareHistory 创建病例接口
func (l *CreateCareHistoryLogic) CreateCareHistory(in *doctor.CreateCareHistoryReq) (*doctor.CreateCareHistoryResp, error) {
	// 1. 参数验证
	if err := l.validateCareHistoryRequest(in); err != nil {
		return &doctor.CreateCareHistoryResp{
			Code:    400,
			Message: err.Error(),
		}, nil
	}

	// 2. 检查病例是否已存在
	var existingCareHistory mysql.HisCareHistory
	if err := l.svcCtx.DB.Where("ch_id = ?", in.ChId).First(&existingCareHistory).Error; err == nil {
		return &doctor.CreateCareHistoryResp{
			Code:    400,
			Message: "病例ID已存在",
		}, nil
	}

	// 3. 检查患者是否存在
	var patient mysql.HisPatient
	if err := l.svcCtx.DB.Where("patient_id = ?", in.PatientId).First(&patient).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &doctor.CreateCareHistoryResp{
				Code:    404,
				Message: "患者不存在",
			}, nil
		}
		return &doctor.CreateCareHistoryResp{
			Code:    500,
			Message: "查询患者信息失败",
		}, nil
	}

	// 4. 检查医生是否存在
	var doctors mysql.SysUser
	if err := l.svcCtx.DB.Where("user_id = ? AND user_type = 'DOCTOR'", in.UserId).First(&doctors).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &doctor.CreateCareHistoryResp{
				Code:    404,
				Message: "医生不存在",
			}, nil
		}
		return &doctor.CreateCareHistoryResp{
			Code:    500,
			Message: "查询医生信息失败",
		}, nil
	}

	// 5. 使用事务创建病例
	tx := l.svcCtx.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 6. 创建病例记录
	careHistory := &mysql.HisCareHistory{
		ChID:         in.ChId,
		UserID:       in.UserId,
		UserName:     in.UserName,
		PatientID:    in.PatientId,
		PatientName:  in.PatientName,
		DeptID:       in.DeptId,
		DeptName:     in.DeptName,
		ReceiveType:  in.ReceiveType,
		IsContagious: in.IsContagious,
		CareTime:     &time.Time{},
		CaseDate:     in.CaseDate,
		RegID:        in.RegId,
		CaseTitle:    in.CaseTitle,
		CaseResult:   in.CaseResult,
		DoctorTips:   in.DoctorTips,
		Remark:       in.Remark,
	}

	if err := tx.Create(careHistory).Error; err != nil {
		tx.Rollback()
		return &doctor.CreateCareHistoryResp{
			Code:    500,
			Message: "创建病例失败",
		}, nil
	}

	// 7. 更新患者最后就诊时间
	if err := tx.Model(&mysql.HisPatient{}).
		Where("patient_id = ?", in.PatientId).
		Update("last_login_time", time.Now()).Error; err != nil {
		tx.Rollback()
		return &doctor.CreateCareHistoryResp{
			Code:    500,
			Message: "更新患者就诊时间失败",
		}, nil
	}

	// 8. 提交事务
	if err := tx.Commit().Error; err != nil {
		return &doctor.CreateCareHistoryResp{
			Code:    500,
			Message: "提交事务失败",
		}, nil
	}

	// 9. 记录操作日志
	go l.recordOperationLog("CREATE_CARE_HISTORY", in.ChId, in.UserName)

	// 10. 发送病例创建通知（异步）
	go l.sendCareHistoryNotification(in.ChId, in.PatientId, in.PatientName)

	return &doctor.CreateCareHistoryResp{
		Code:    200,
		Message: "创建病例成功",
		ChId:    in.ChId,
	}, nil
}

// validateCareHistoryRequest 验证病例请求参数
func (l *CreateCareHistoryLogic) validateCareHistoryRequest(in *doctor.CreateCareHistoryReq) error {
	if in.ChId == "" {
		return fmt.Errorf("病例ID不能为空")
	}
	if in.PatientId == "" {
		return fmt.Errorf("患者ID不能为空")
	}
	if in.PatientName == "" {
		return fmt.Errorf("患者姓名不能为空")
	}
	if in.UserId == 0 {
		return fmt.Errorf("医生ID不能为空")
	}
	if in.UserName == "" {
		return fmt.Errorf("医生姓名不能为空")
	}
	if in.DeptId == 0 {
		return fmt.Errorf("科室ID不能为空")
	}
	if in.DeptName == "" {
		return fmt.Errorf("科室名称不能为空")
	}
	if in.CaseDate == "" {
		return fmt.Errorf("病例日期不能为空")
	}
	if in.CaseTitle == "" {
		return fmt.Errorf("病例标题不能为空")
	}
	return nil
}

// recordOperationLog 记录操作日志
func (l *CreateCareHistoryLogic) recordOperationLog(operation, targetId, operator string) {
	log := &mysql.SysOperLog{
		Title:        "创建病例",
		BusinessType: operation,
		Method:       "CreateCareHistory",
		OperName:     operator,
		OperURL:      "/doctor/carehistory/create",
		OperParam:    fmt.Sprintf("病例ID: %s", targetId),
		Status:       "0", // 成功
		OperTime:     &time.Time{},
	}

	if err := l.svcCtx.DB.Create(log).Error; err != nil {
		l.Logger.Errorf("记录操作日志失败: %v", err)
	}
}

// sendCareHistoryNotification 发送病例创建通知
func (l *CreateCareHistoryLogic) sendCareHistoryNotification(chId, patientId, patientName string) {
	// 这里可以实现发送短信、邮件或推送通知的逻辑
	// 例如：发送短信通知患者病例已创建
	notificationType := "SMS"
	content := fmt.Sprintf("患者 %s 的病例 %s 已创建，请及时查看。", patientName, chId)
	fmt.Println(content)
	// 记录短信发送日志
	smsLog := &mysql.SysSmsLog{
		ID:         time.Now().UnixNano(),
		Mobile:     "", // 需要从患者信息中获取手机号
		CreateTime: &time.Time{},
		Code:       "CARE_HISTORY_CREATED",
		Status:     "SENT",
		Type:       notificationType,
	}

	if err := l.svcCtx.DB.Create(smsLog).Error; err != nil {
		l.Logger.Errorf("记录短信发送日志失败: %v", err)
	}

	l.Logger.Infof("发送病例创建通知: 患者ID=%s, 病例ID=%s", patientId, chId)
}
