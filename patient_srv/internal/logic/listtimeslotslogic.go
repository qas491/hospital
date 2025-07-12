package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/qas491/hospital/patient_srv/model/mysql"

	"github.com/qas491/hospital/patient_srv/internal/svc"
	"github.com/qas491/hospital/patient_srv/patient"

	"github.com/zeromicro/go-zero/core/logx"
)

// ListTimeSlotsLogic 时间段列表查询逻辑处理器
// 负责根据医生ID和日期查询可预约的时间段信息
type ListTimeSlotsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewListTimeSlotsLogic 创建时间段列表查询逻辑处理器实例
// @param ctx 上下文信息
// @param svcCtx 服务上下文，包含数据库连接等资源
// @return *ListTimeSlotsLogic 时间段列表查询逻辑处理器实例
func NewListTimeSlotsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTimeSlotsLogic {
	return &ListTimeSlotsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// TimeSlot 时间段信息数据模型
// 对应数据库表 his_scheduling，用于存储医生排班信息
type TimeSlot struct {
	DoctorID       int32      `gorm:"column:user_id"`         // 医生ID
	DeptID         int32      `gorm:"column:dept_id"`         // 科室ID
	Date           string     `gorm:"column:scheduling_day"`  // 排班日期
	SubsectionType string     `gorm:"column:subsection_type"` // 时间段类型（上午/下午）
	SchedulingType string     `gorm:"column:scheduling_type"` // 排班类型
	CreateTime     *time.Time `gorm:"column:create_time"`     // 创建时间
	CreateBy       string     `gorm:"column:create_by"`       // 创建人
}

// TableName 指定数据库表名
// @return string 数据库表名
func (TimeSlot) TableName() string {
	return "his_scheduling"
}

// subsectionToTime 将时间段类型转换为具体的时间范围
// @param subsection 时间段类型（morning/afternoon）
// @return string 开始时间
// @return string 结束时间
func subsectionToTime(subsection string) (string, string) {
	switch subsection {
	case "morning":
		return "08:00", "12:00"
	case "afternoon":
		return "13:00", "17:00"
	default:
		return "00:00", "00:00"
	}
}

// ListTimeSlots 查询时间段列表
// 根据医生ID和日期查询该医生的可预约时间段，并转换为proto响应格式
// @param req 时间段列表查询请求，包含医生ID和日期
// @return *seckill.ListTimeSlotsResponse 时间段列表响应
// @return error 错误信息
func (l *ListTimeSlotsLogic) ListTimeSlots(req *patient.ListTimeSlotsRequest) (*patient.ListTimeSlotsResponse, error) {
	// 参数验证
	if req.DoctorId <= 0 {
		return nil, fmt.Errorf("医生ID无效")
	}
	if req.Date == "" {
		return nil, fmt.Errorf("日期不能为空")
	}

	// 检查数据库连接
	db, err := mysql.GetDB()
	if err != nil {
		l.Logger.Errorf("数据库连接失败: %v", err)
		return nil, fmt.Errorf("数据库连接失败")
	}

	// 根据医生ID和日期查询排班信息
	var slots []TimeSlot
	if err := db.Where("user_id = ? AND scheduling_day = ?", req.DoctorId, req.Date).Find(&slots).Error; err != nil {
		l.Logger.Errorf("查询医生 %d 在 %s 的排班信息失败: %v", req.DoctorId, req.Date, err)
		return nil, fmt.Errorf("查询排班信息失败: %v", err)
	}

	// 构建响应数据
	resp := &patient.ListTimeSlotsResponse{}
	for i, s := range slots {
		// 将时间段类型转换为具体时间
		start, end := subsectionToTime(s.SubsectionType)

		// 计算可用名额（实际应该查询已预约数量）
		available := l.calculateAvailableSlots(req.DoctorId, req.Date, s.SubsectionType)

		// 生成时间段ID（使用索引作为ID）
		slotId := int32(i + 1)

		// 将数据库模型转换为proto响应格式
		resp.Timeslots = append(resp.Timeslots, &patient.TimeSlot{
			Id:        slotId,
			DoctorId:  s.DoctorID,
			Date:      s.Date,
			StartTime: start,
			EndTime:   end,
			Available: available,
		})
	}

	l.Logger.Infof("成功查询到医生 %d 在 %s 的 %d 个时间段", req.DoctorId, req.Date, len(resp.Timeslots))
	return resp, nil
}

// calculateAvailableSlots 计算可用名额
// @param doctorId 医生ID
// @param date 日期
// @param subsectionType 时间段类型
// @return int32 可用名额
func (l *ListTimeSlotsLogic) calculateAvailableSlots(doctorId int32, date string, subsectionType string) int32 {
	// 这里应该查询已预约的数量，然后计算剩余名额
	// 暂时返回固定值，实际应该从数据库查询
	var totalSlots int32 = 20 // 每个时间段总名额
	var bookedSlots int32 = 0 // 已预约名额

	// TODO: 实现从数据库查询已预约数量的逻辑
	// 示例查询：
	// SELECT COUNT(*) FROM his_registration
	// WHERE user_id = ? AND visit_date = ? AND subsection_type = ? AND registration_status != 'cancelled'

	return totalSlots - bookedSlots
}
