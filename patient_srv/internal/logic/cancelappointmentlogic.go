package logic

import (
	"context"
	"fmt"
	"github.com/qas491/hospital/patient_srv/model/mysql"

	"github.com/qas491/hospital/patient_srv/internal/svc"
	"github.com/qas491/hospital/patient_srv/patient"

	"github.com/zeromicro/go-zero/core/logx"
)

// CancelAppointmentLogic 取消预约逻辑处理器
// 负责处理患者取消预约的业务逻辑
type CancelAppointmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewCancelAppointmentLogic 创建取消预约逻辑处理器实例
// @param ctx 上下文信息
// @param svcCtx 服务上下文，包含数据库连接等资源
// @return *CancelAppointmentLogic 取消预约逻辑处理器实例
func NewCancelAppointmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelAppointmentLogic {
	return &CancelAppointmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CancelAppointment 取消预约
// 根据预约ID更新挂号记录状态为已取消
// @param req 取消预约请求，包含预约ID
// @return *seckill.CancelAppointmentResponse 取消预约响应
// @return error 错误信息

func (l *CancelAppointmentLogic) CancelAppointment(req *patient.CancelAppointmentRequest) (*patient.CancelAppointmentResponse, error) {
	// 检查数据库连接
	db, err := mysql.GetDB()
	if err != nil {
		l.Logger.Errorf("数据库连接失败: %v", err)
		return nil, fmt.Errorf("数据库连接失败")
	}

	// 更新挂号记录状态为已取消
	if err := db.Model(&mysql.HisRegistration{}).
		Where("registration_id = ?", req.AppointmentId).
		Update("registration_status", "cancelled").Error; err != nil {
		l.Logger.Errorf("取消预约ID %s 失败: %v", req.AppointmentId, err)
		return &patient.CancelAppointmentResponse{Success: false}, err
	}

	l.Logger.Infof("成功取消预约，预约ID: %s", req.AppointmentId)
	return &patient.CancelAppointmentResponse{Success: true}, nil
}
