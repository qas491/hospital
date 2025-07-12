package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCareHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCareHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCareHistoryListLogic {
	return &GetCareHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCareHistoryListLogic) GetCareHistoryList(req *types.GetCareHistoryListReq) (*types.GetCareHistoryListResp, error) {
	rpcResp, err := l.svcCtx.DoctorRpc.GetCareHistoryList(l.ctx, &doctor.GetCareHistoryListReq{
		Page:        req.Page,
		PageSize:    req.Page_size,
		PatientId:   req.Patient_id,
		PatientName: req.Patient_name,
		UserId:      req.User_id,
		DeptId:      req.Dept_id,
		CaseDate:    req.Case_date,
		StartTime:   req.Start_time,
		EndTime:     req.End_time,
	})
	if err != nil {
		return nil, err
	}
	var list []types.CareHistoryInfo
	for _, c := range rpcResp.List {
		list = append(list, convertToCareHistoryInfo(c))
	}
	return &types.GetCareHistoryListResp{
		Code:    rpcResp.Code,
		Message: rpcResp.Message,
		List:    list,
		Total:   rpcResp.Total,
	}, nil
}
