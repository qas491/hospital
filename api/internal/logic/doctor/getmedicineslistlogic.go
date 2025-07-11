package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMedicinesListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMedicinesListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMedicinesListLogic {
	return &GetMedicinesListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMedicinesListLogic) GetMedicinesList(req *types.GetMedicinesListReq) (resp *types.GetMedicinesListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
