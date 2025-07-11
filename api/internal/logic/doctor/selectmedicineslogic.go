package doctor

import (
	"context"

	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectMedicinesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectMedicinesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectMedicinesLogic {
	return &SelectMedicinesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectMedicinesLogic) SelectMedicines(req *types.SelectMedicinesReq) (resp *types.SelectMedicinesResp, err error) {
	// todo: add your logic here and delete this line

	return
}
