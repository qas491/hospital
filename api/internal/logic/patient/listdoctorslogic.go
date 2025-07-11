package patient

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

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

func (l *ListDoctorsLogic) ListDoctors() (resp *types.ListDoctorsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
