package patient

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"
	"github.com/qas491/hospital/patient_srv/patient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDepartmentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListDepartmentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDepartmentsLogic {
	return &ListDepartmentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ListDepartments 获取科室列表
func (l *ListDepartmentsLogic) ListDepartments(req *types.ListDepartmentsReq) (*types.ListDepartmentsResp, error) {
	// 调用 patientRpc 的 ListDepartments 方法
	rpcResp, err := l.svcCtx.PatientRpc.ListDepartments(l.ctx, &patient.ListDepartmentsRequest{})
	if err != nil {
		return nil, err // RPC 调用失败，返回错误
	}

	// 类型转换，将 RPC 返回的科室列表转换为 API 层结构体
	var departments []types.Department
	for _, d := range rpcResp.Departments {
		departments = append(departments, types.Department{
			Id:          int(d.Id), // int32 转 int
			Name:        d.Name,
			Description: d.Description,
		})
	}

	// 返回 API 层定义的响应结构体
	return &types.ListDepartmentsResp{
		Departments: departments,
	}, nil
}
