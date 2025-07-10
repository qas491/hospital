package logic

import (
	"context"
	"fmt"
	"github.com/qas491/hospital/patient_srv/model/mysql"
	"strconv"

	"github.com/qas491/hospital/patient_srv/internal/svc"
	"github.com/qas491/hospital/patient_srv/patient"

	"github.com/zeromicro/go-zero/core/logx"
)

// GetAppointmentLogic 查询预约信息逻辑处理器
// 负责根据预约ID查询预约详细信息
type GetAppointmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewGetAppointmentLogic 创建查询预约信息逻辑处理器实例
// @param ctx 上下文信息
// @param svcCtx 服务上下文，包含数据库连接等资源
// @return *GetAppointmentLogic 查询预约信息逻辑处理器实例
func NewGetAppointmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppointmentLogic {
	return &GetAppointmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetAppointment 查询预约信息
// 根据预约ID查询挂号记录的详细信息，并转换为proto响应格式
// @param req 查询预约请求，包含预约ID
// @return *seckill.GetAppointmentResponse 预约信息响应
// @return error 错误信息

func (l *GetAppointmentLogic) GetAppointment(req *patient.GetAppointmentRequest) (*patient.GetAppointmentResponse, error) {
	// 检查数据库连接
	db, err := mysql.GetDB()
	if err != nil {
		l.Logger.Errorf("数据库连接失败: %v", err)
		return nil, fmt.Errorf("数据库连接失败")
	}

	// 根据预约ID查询挂号记录
	var reg mysql.HisRegistration
	if err := db.Where("registration_id = ?", req.AppointmentId).First(&reg).Error; err != nil {
		l.Logger.Errorf("查询预约ID %s 的挂号记录失败: %v", req.AppointmentId, err)
		return nil, err
	}

	// 将患者ID从字符串转换为整数
	patientId, _ := strconv.Atoi(reg.PatientID)

	// 构建响应数据
	resp := &patient.GetAppointmentResponse{
		Appointment: &patient.Appointment{
			Id:           0, // 可用自增ID或RegistrationID
			PatientId:    int32(patientId),
			DoctorId:     int32(reg.UserID),
			DepartmentId: int32(reg.DeptID),
			TimeslotId:   0, // 需补充，可根据业务需求添加时间段ID
			Status:       reg.RegistrationStatus,
			CreatedAt:    reg.CreateTime.Format("2006-01-02 15:04:05"),
		},
	}

	l.Logger.Infof("成功查询预约信息，预约ID: %s, 患者: %s, 医生: %s",
		req.AppointmentId, reg.PatientName, reg.DoctorName)
	return resp, nil
}
