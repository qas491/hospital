package logic

import (
	"context"
	"fmt"
	"github.com/qas491/hospital/patient_srv/model/mysql"

	"github.com/qas491/hospital/patient_srv/internal/svc"
	"github.com/qas491/hospital/patient_srv/patient"

	"github.com/zeromicro/go-zero/core/logx"
)

// ListDoctorsLogic 医生列表查询逻辑处理器
// 负责根据科室ID查询该科室下的所有医生信息
type ListDoctorsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewListDoctorsLogic 创建医生列表查询逻辑处理器实例
// @param ctx 上下文信息
// @param svcCtx 服务上下文，包含数据库连接等资源
// @return *ListDoctorsLogic 医生列表查询逻辑处理器实例
func NewListDoctorsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDoctorsLogic {
	return &ListDoctorsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Doctor 医生信息数据模型
// 对应数据库表 sys_user，用于存储医生基本信息
type Doctor struct {
	ID           int32  `gorm:"column:user_id;primaryKey" json:"id"` // 医生ID，主键
	Name         string `gorm:"column:user_name" json:"name"`        // 医生姓名
	DepartmentID int32  `gorm:"column:dept_id" json:"department_id"` // 所属科室ID
	Title        string `gorm:"column:user_rank" json:"title"`       // 医生职称
	Profile      string `gorm:"column:introduction" json:"profile"`  // 医生简介
}

// TableName 指定数据库表名
// @return string 数据库表名
func (Doctor) TableName() string {
	return "sys_user"
}

// ListDoctors 查询医生列表
// 根据科室ID查询该科室下的所有医生信息，并转换为proto响应格式
// @param req 医生列表查询请求，包含科室ID
// @return *seckill.ListDoctorsResponse 医生列表响应
// @return error 错误信息
func (l *ListDoctorsLogic) ListDoctors(req *patient.ListDoctorsRequest) (*patient.ListDoctorsResponse, error) {
	// 检查数据库连接
	db, err := mysql.GetDB()
	if err != nil {
		l.Logger.Errorf("数据库连接失败: %v", err)
		return nil, fmt.Errorf("数据库连接失败")
	}

	// 根据科室ID查询医生信息
	var doctors []Doctor
	if err := db.Where("dept_id = ?", req.DepartmentId).Find(&doctors).Error; err != nil {
		l.Logger.Errorf("查询科室 %d 的医生列表失败: %v", req.DepartmentId, err)
		return nil, err
	}

	// 构建响应数据
	resp := &patient.ListDoctorsResponse{}
	for _, d := range doctors {
		// 将数据库模型转换为proto响应格式
		resp.Doctors = append(resp.Doctors, &patient.Doctor{
			Id:           d.ID,
			Name:         d.Name,
			DepartmentId: d.DepartmentID,
			Title:        d.Title,
			Profile:      d.Profile,
		})
	}

	l.Logger.Infof("成功查询到科室 %d 的 %d 位医生", req.DepartmentId, len(resp.Doctors))
	return resp, nil
}
