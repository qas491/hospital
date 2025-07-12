package patient

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"
	"github.com/qas491/hospital/patient_srv/patient"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListTimeSlotsLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	DoctorId int32 // 新增
}

func NewListTimeSlotsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTimeSlotsLogic {
	return &ListTimeSlotsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTimeSlotsLogic) ListTimeSlots(req *types.ListTimeSlotsReq) (*types.ListTimeSlotsResp, error) {
	rpcResp, err := l.svcCtx.PatientRpc.ListTimeSlots(l.ctx, &patient.ListTimeSlotsRequest{
		DoctorId: l.DoctorId, // 用 logic 结构体的 DoctorId 字段
		Date:     req.Date,
	})
	if err != nil {
		return nil, err
	}
	var timeslots []types.TimeSlot
	for _, t := range rpcResp.Timeslots {
		timeslots = append(timeslots, types.TimeSlot{
			Id:         t.Id,
			Doctor_id:  t.DoctorId,
			Date:       t.Date,
			Start_time: t.StartTime,
			End_time:   t.EndTime,
			Available:  t.Available,
		})
	}
	return &types.ListTimeSlotsResp{
		Timeslots: timeslots,
	}, nil
}
