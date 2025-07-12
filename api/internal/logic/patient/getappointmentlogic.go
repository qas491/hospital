package patient

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"
	"github.com/qas491/hospital/patient_srv/patient"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppointmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAppointmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppointmentLogic {
	return &GetAppointmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 实现 GetAppointment 方法
func (l *GetAppointmentLogic) GetAppointment(req *types.GetAppointmentReq) (*types.GetAppointmentResp, error) {
	// 调用 patientRpc 的 GetAppointment 方法，传递预约ID
	rpcResp, err := l.svcCtx.PatientRpc.GetAppointment(l.ctx, &patient.GetAppointmentRequest{
		AppointmentId: req.Appointment_id, // 预约ID
	})
	if err != nil {
		return nil, err // RPC 调用失败，返回错误
	}
	// 类型转换
	appt := rpcResp.Appointment
	return &types.GetAppointmentResp{
		Appointment: types.Appointment{
			Id:            appt.Id,
			Patient_id:    appt.PatientId,
			Doctor_id:     appt.DoctorId,
			Department_id: appt.DepartmentId,
			Timeslot_id:   appt.TimeslotId,
			Status:        appt.Status,
			Created_at:    appt.CreatedAt,
		},
	}, nil
}
