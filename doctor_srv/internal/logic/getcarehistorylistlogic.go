package logic

import (
	"context"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/qas491/hospital/doctor_srv/internal/svc"
	"github.com/qas491/hospital/doctor_srv/model/mysql"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCareHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCareHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCareHistoryListLogic {
	return &GetCareHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取病例列表
func (l *GetCareHistoryListLogic) GetCareHistoryList(in *doctor.GetCareHistoryListReq) (*doctor.GetCareHistoryListResp, error) {
	db := l.svcCtx.DB.Model(&mysql.HisCareHistory{})
	if in.PatientId != "" {
		db = db.Where("patient_id = ?", in.PatientId)
	}
	if in.PatientName != "" {
		db = db.Where("patient_name LIKE ?", "%"+in.PatientName+"%")
	}
	if in.UserId != 0 {
		db = db.Where("user_id = ?", in.UserId)
	}
	if in.DeptId != "" {
		db = db.Where("dept_id = ?", in.DeptId)
	}
	if in.CaseDate != "" {
		db = db.Where("case_date = ?", in.CaseDate)
	}
	if in.StartTime != "" {
		db = db.Where("care_time >= ?", in.StartTime)
	}
	if in.EndTime != "" {
		db = db.Where("care_time <= ?", in.EndTime)
	}
	var total int64
	db.Count(&total)
	var list []mysql.HisCareHistory
	db = db.Order("care_time desc").Offset(int((in.Page - 1) * in.PageSize)).Limit(int(in.PageSize))
	db.Find(&list)
	respList := make([]*doctor.CareHistoryInfo, 0, len(list))
	for _, h := range list {
		respList = append(respList, &doctor.CareHistoryInfo{
			ChId:         h.ChID,
			UserId:       h.UserID,
			UserName:     h.UserName,
			PatientId:    h.PatientID,
			PatientName:  h.PatientName,
			DeptId:       h.DeptID,
			DeptName:     h.DeptName,
			ReceiveType:  h.ReceiveType,
			IsContagious: h.IsContagious,
			CareTime:     FormatTime(h.CareTime),
			CaseDate:     h.CaseDate,
			RegId:        h.RegID,
			CaseTitle:    h.CaseTitle,
			CaseResult:   h.CaseResult,
			DoctorTips:   h.DoctorTips,
			Remark:       h.Remark,
		})
	}
	return &doctor.GetCareHistoryListResp{
		Code:    0,
		Message: "ok",
		List:    respList,
		Total:   total,
	}, nil
}
