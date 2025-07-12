package patient

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"
	"github.com/qas491/hospital/patient_srv/patient"

	"github.com/zeromicro/go-zero/core/logx"
)

// CancelAppointmentLogic 处理取消预约的业务逻辑
type CancelAppointmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelAppointmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelAppointmentLogic {
	return &CancelAppointmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelAppointmentLogic) CancelAppointment(req *types.CancelAppointmentReq) (*types.CancelAppointmentResp, error) {
	// 调用 patientRpc 的 CancelAppointment 方法，传递预约ID
	rpcResp, err := l.svcCtx.PatientRpc.CancelAppointment(l.ctx, &patient.CancelAppointmentRequest{
		AppointmentId: req.Appointment_id, // 预约ID
	})
	if err != nil {
		return nil, err // RPC 调用失败，返回错误
	}
	// 返回 API 层定义的响应结构体
	return &types.CancelAppointmentResp{
		Success: rpcResp.Success, // 是否成功
	}, nil
}
