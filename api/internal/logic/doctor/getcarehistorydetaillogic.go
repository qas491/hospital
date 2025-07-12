package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCareHistoryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCareHistoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCareHistoryDetailLogic {
	return &GetCareHistoryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCareHistoryDetailLogic) GetCareHistoryDetail(req *types.GetCareHistoryDetailReq) (*types.GetCareHistoryDetailResp, error) {
	rpcResp, err := l.svcCtx.DoctorRpc.GetCareHistoryDetail(l.ctx, &doctor.GetCareHistoryDetailReq{
		ChId: req.Ch_id,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetCareHistoryDetailResp{
		Code:    rpcResp.Code,
		Message: rpcResp.Message,
		Detail:  convertToCareHistoryInfo(rpcResp.Detail),
	}, nil
}

func convertToCareHistoryInfo(in *doctor.CareHistoryInfo) types.CareHistoryInfo {
	if in == nil {
		return types.CareHistoryInfo{}
	}
	return types.CareHistoryInfo{
		Ch_id:         in.ChId,
		User_id:       in.UserId,
		User_name:     in.UserName,
		Patient_id:    in.PatientId,
		Patient_name:  in.PatientName,
		Dept_id:       in.DeptId,
		Dept_name:     in.DeptName,
		Receive_type:  in.ReceiveType,
		Is_contagious: in.IsContagious,
		Care_time:     in.CareTime,
		Case_date:     in.CaseDate,
		Reg_id:        in.RegId,
		Case_title:    in.CaseTitle,
		Case_result:   in.CaseResult,
		Doctor_tips:   in.DoctorTips,
		Remark:        in.Remark,
	}
}
