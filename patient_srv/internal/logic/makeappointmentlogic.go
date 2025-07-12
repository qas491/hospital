package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/qas491/hospital/patient_srv/model/mysql"

	"github.com/qas491/hospital/patient_srv/internal/svc"
	"github.com/qas491/hospital/patient_srv/patient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MakeAppointmentLogic 预约挂号逻辑处理器
// 负责处理患者预约挂号的核心业务逻辑
type MakeAppointmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewMakeAppointmentLogic 创建预约挂号逻辑处理器实例
// @param ctx 上下文信息
// @param svcCtx 服务上下文，包含数据库连接等资源
// @return *MakeAppointmentLogic 预约挂号逻辑处理器实例
func NewMakeAppointmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakeAppointmentLogic {
	return &MakeAppointmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MakeAppointment 创建预约挂号
// 验证患者和医生信息，创建挂号记录，并返回预约结果
// @param req 预约挂号请求，包含患者ID、医生ID、科室ID、时间段ID
// @return *seckill.MakeAppointmentResponse 预约挂号响应
// @return error 错误信息
func (l *MakeAppointmentLogic) MakeAppointment(req *patient.MakeAppointmentRequest) (*patient.MakeAppointmentResponse, error) {
	// 检查数据库连接
	db, err := mysql.GetDB()
	if err != nil {
		l.Logger.Errorf("数据库连接失败: %v", err)
		return nil, fmt.Errorf("数据库连接失败")
	}

	// 打印收到的 PatientId
	l.Logger.Infof("收到预约请求，req: %+v", req)

	// 1. 验证患者信息是否存在
	var patients mysql.HisPatient
	if err := db.Where("patient_id = ?", req.PatientId).First(&patients).Error; err != nil {
		l.Logger.Errorf("患者 %d 不存在: %v", req.PatientId, err)
		return nil, fmt.Errorf("患者不存在")
	}

	// 2. 验证医生信息是否存在
	var doctor mysql.SysUser
	if err := db.Where("user_id = ?", req.DoctorId).First(&doctor).Error; err != nil {
		l.Logger.Errorf("医生 %d 不存在: %v", req.DoctorId, err)
		return nil, fmt.Errorf("医生不存在")
	}

	// 3. 创建挂号记录
	reg := mysql.HisRegistration{
		RegistrationID:     uuid.New().String(),             // 生成唯一挂号ID
		PatientID:          patients.PatientID,              // 患者ID
		PatientName:        patients.Name,                   // 患者姓名
		UserID:             int64(req.DoctorId),             // 医生ID
		DoctorName:         doctor.UserName,                 // 医生姓名
		DeptID:             int64(req.DepartmentId),         // 科室ID
		RegistrationStatus: "pending",                       // 预约状态：待确认
		VisitDate:          time.Now().Format("2006-01-02"), // 就诊日期
		CreateTime:         timePtr(time.Now()),             // 创建时间
	}

	// 4. 保存挂号记录到数据库
	if err := db.Create(&reg).Error; err != nil {
		l.Logger.Errorf("创建挂号记录失败: %v", err)
		return nil, err
	}

	// 5. 构建响应数据
	resp := &patient.MakeAppointmentResponse{
		Appointment: &patient.Appointment{
			Id:           1, // 可用自增ID或RegistrationID
			PatientId:    req.PatientId,
			DoctorId:     req.DoctorId,
			DepartmentId: req.DepartmentId,
			TimeslotId:   req.TimeslotId,
			Status:       "pending",
			CreatedAt:    reg.CreateTime.Format("2006-01-02 15:04:05"),
		},
	}

	l.Logger.Infof("成功创建预约挂号，患者: %s, 医生: %s, 挂号ID: %s",
		patients.Name, doctor.UserName, reg.RegistrationID)
	return resp, nil
}

// timePtr 返回时间指针
// @param t 时间值
// @return *time.Time 时间指针
func timePtr(t time.Time) *time.Time {
	return &t
}
