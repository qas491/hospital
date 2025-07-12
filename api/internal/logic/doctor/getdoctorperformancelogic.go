package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDoctorPerformanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDoctorPerformanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDoctorPerformanceLogic {
	return &GetDoctorPerformanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDoctorPerformanceLogic) GetDoctorPerformance(req *types.GetDoctorPerformanceReq) (*types.GetDoctorPerformanceResp, error) {
	rpcResp, err := l.svcCtx.DoctorRpc.GetDoctorPerformance(l.ctx, &doctor.GetDoctorPerformanceReq{
		DoctorId:  req.Doctor_id,
		StartDate: req.Start_date,
		EndDate:   req.End_date,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetDoctorPerformanceResp{
		Code:        rpcResp.Code,
		Message:     rpcResp.Message,
		Performance: convertToDoctorPerformanceInfo(rpcResp.Performance),
	}, nil
}

func convertToDoctorPerformanceInfo(in *doctor.DoctorPerformanceInfo) types.DoctorPerformanceInfo {
	if in == nil {
		return types.DoctorPerformanceInfo{}
	}
	var details []types.PerformanceDetail
	for _, d := range in.PerformanceDetails {
		details = append(details, types.PerformanceDetail{
			Performance_id:     d.PerformanceId,
			Co_id:              d.CoId,
			Payment_amount:     d.PaymentAmount,
			Performance_amount: d.PerformanceAmount,
			Performance_rate:   d.PerformanceRate,
			Performance_date:   d.PerformanceDate,
		})
	}
	return types.DoctorPerformanceInfo{
		Doctor_id:           in.DoctorId,
		Doctor_name:         in.DoctorName,
		Dept_name:           in.DeptName,
		Total_performance:   in.TotalPerformance,
		Prescription_count:  in.PrescriptionCount,
		Payment_count:       in.PaymentCount,
		Performance_details: details,
	}
}
