package patient

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"
	"github.com/qas491/hospital/patient_srv/patient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDoctorsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListDoctorsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDoctorsLogic {
	return &ListDoctorsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ListDoctors 获取指定科室的医生列表
func (l *ListDoctorsLogic) ListDoctors(req *types.ListDoctorsReq) (*types.ListDoctorsResp, error) {
	// 调用 patientRpc 的 ListDoctors 方法，传递科室ID
	rpcResp, err := l.svcCtx.PatientRpc.ListDoctors(l.ctx, &patient.ListDoctorsRequest{
		DepartmentId: req.Department_id, // 科室ID
	})
	if err != nil {
		return nil, err // RPC 调用失败，返回错误
	}

	// 类型转换，将 RPC 返回的医生列表转换为 API 层结构体
	var doctors []types.Doctor
	for _, d := range rpcResp.Doctors {
		doctors = append(doctors, types.Doctor{
			Id:            d.Id,
			Name:          d.Name,
			Department_id: d.DepartmentId,
			Title:         d.Title,
			Profile:       d.Profile,
		})
	}

	// 返回 API 层定义的响应结构体
	return &types.ListDoctorsResp{
		Doctors: doctors,
	}, nil
}
