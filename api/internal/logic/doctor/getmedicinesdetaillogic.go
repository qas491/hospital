package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMedicinesDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMedicinesDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMedicinesDetailLogic {
	return &GetMedicinesDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMedicinesDetailLogic) GetMedicinesDetail(req *types.GetMedicinesDetailReq) (resp *types.GetMedicinesDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
