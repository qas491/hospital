package patient

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"
	"github.com/qas491/hospital/patient_srv/patient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MakeAppointmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMakeAppointmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakeAppointmentLogic {
	return &MakeAppointmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MakeAppointmentLogic) MakeAppointment(req *types.MakeAppointmentReq) (*types.MakeAppointmentResp, error) {
	rpcResp, err := l.svcCtx.PatientRpc.MakeAppointment(l.ctx, &patient.MakeAppointmentRequest{
		PatientId:    req.Patient_id,
		DoctorId:     req.Doctor_id,
		DepartmentId: req.Department_id,
		TimeslotId:   req.Timeslot_id,
	})
	if err != nil {
		return nil, err
	}
	return &types.MakeAppointmentResp{
		Appointment: types.Appointment{
			Id:            rpcResp.Appointment.Id,
			Patient_id:    rpcResp.Appointment.PatientId,
			Doctor_id:     rpcResp.Appointment.DoctorId,
			Department_id: rpcResp.Appointment.DepartmentId,
			Timeslot_id:   rpcResp.Appointment.TimeslotId,
			Status:        rpcResp.Appointment.Status,
			Created_at:    rpcResp.Appointment.CreatedAt,
		},
	}, nil
}
