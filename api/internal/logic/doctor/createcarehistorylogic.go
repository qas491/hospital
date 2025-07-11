package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

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

func (l *CreateCareHistoryLogic) CreateCareHistory(req *types.CreateCareHistoryReq) (resp *types.CreateCareHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
