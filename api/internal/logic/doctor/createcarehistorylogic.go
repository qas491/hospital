package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCareHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCareHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCareHistoryLogic {
	return &CreateCareHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCareHistoryLogic) CreateCareHistory(req *types.CreateCareHistoryReq) (*types.CreateCareHistoryResp, error) {
	rpcResp, err := l.svcCtx.DoctorRpc.CreateCareHistory(l.ctx, &doctor.CreateCareHistoryReq{
		ChId:         req.Ch_id,
		UserId:       req.User_id,
		UserName:     req.User_name,
		PatientId:    req.Patient_id,
		PatientName:  req.Patient_name,
		DeptId:       req.Dept_id,
		DeptName:     req.Dept_name,
		ReceiveType:  req.Receive_type,
		IsContagious: req.Is_contagious,
		CaseDate:     req.Case_date,
		RegId:        req.Reg_id,
		CaseTitle:    req.Case_title,
		CaseResult:   req.Case_result,
		DoctorTips:   req.Doctor_tips,
		Remark:       req.Remark,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateCareHistoryResp{
		Code:    rpcResp.Code,
		Message: rpcResp.Message,
		Ch_id:   rpcResp.ChId,
	}, nil
}
