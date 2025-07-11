package logic

import (
	"context"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/qas491/hospital/doctor_srv/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDoctorPerformanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDoctorPerformanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDoctorPerformanceLogic {
	return &GetDoctorPerformanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取医生业绩
func (l *GetDoctorPerformanceLogic) GetDoctorPerformance(in *doctor.GetDoctorPerformanceReq) (*doctor.GetDoctorPerformanceResp, error) {
	// todo: add your logic here and delete this line

	return &doctor.GetDoctorPerformanceResp{}, nil
}
