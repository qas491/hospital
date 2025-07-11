package logic

import (
	"context"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/qas491/hospital/doctor_srv/internal/svc"
	"github.com/qas491/hospital/doctor_srv/model/mysql"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCareHistoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCareHistoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCareHistoryDetailLogic {
	return &GetCareHistoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取病例详情
func (l *GetCareHistoryDetailLogic) GetCareHistoryDetail(in *doctor.GetCareHistoryDetailReq) (*doctor.GetCareHistoryDetailResp, error) {
	var history mysql.HisCareHistory
	err := l.svcCtx.DB.Where("ch_id = ?", in.ChId).First(&history).Error
	if err != nil {
		return &doctor.GetCareHistoryDetailResp{
			Code:    1,
			Message: "未找到病例: " + err.Error(),
		}, nil
	}
	resp := &doctor.GetCareHistoryDetailResp{
		Code:    0,
		Message: "ok",
		Detail: &doctor.CareHistoryInfo{
			ChId:         history.ChID,
			UserId:       history.UserID,
			UserName:     history.UserName,
			PatientId:    history.PatientID,
			PatientName:  history.PatientName,
			DeptId:       history.DeptID,
			DeptName:     history.DeptName,
			ReceiveType:  history.ReceiveType,
			IsContagious: history.IsContagious,
			CareTime:     FormatTime(history.CareTime),
			CaseDate:     history.CaseDate,
			RegId:        history.RegID,
			CaseTitle:    history.CaseTitle,
			CaseResult:   history.CaseResult,
			DoctorTips:   history.DoctorTips,
			Remark:       history.Remark,
		},
	}
	return resp, nil
}
