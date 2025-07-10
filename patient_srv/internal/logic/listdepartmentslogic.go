package logic

import (
	"context"
	"fmt"
	"github.com/qas491/hospital/patient_srv/model/mysql"

	"github.com/qas491/hospital/patient_srv/internal/svc"
	"github.com/qas491/hospital/patient_srv/patient"

	"github.com/zeromicro/go-zero/core/logx"
)

// ListDepartmentsLogic 科室列表查询逻辑处理器
// 负责查询医院所有科室信息并返回给客户端
type ListDepartmentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewListDepartmentsLogic 创建科室列表查询逻辑处理器实例
// @param ctx 上下文信息
// @param svcCtx 服务上下文，包含数据库连接等资源
// @return *ListDepartmentsLogic 科室列表查询逻辑处理器实例
func NewListDepartmentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDepartmentsLogic {
	return &ListDepartmentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Department 科室信息数据模型
// 对应数据库表 sys_role，用于存储科室基本信息
type Department struct {
	ID          int32  `gorm:"column:role_id;primaryKey" json:"id"` // 科室ID，主键
	Name        string `gorm:"column:role_name" json:"name"`        // 科室名称
	Description string `gorm:"column:remark" json:"description"`    // 科室描述信息
}

// TableName 指定数据库表名
// @return string 数据库表名
func (Department) TableName() string {
	return "sys_role"
}

// ListDepartments 查询科室列表
// 从数据库中查询所有科室信息，并转换为proto响应格式
// @param req 科室列表查询请求
// @return *seckill.ListDepartmentsResponse 科室列表响应
// @return error 错误信息
func (l *ListDepartmentsLogic) ListDepartments(req *patient.ListDepartmentsRequest) (*patient.ListDepartmentsResponse, error) {
	// 检查数据库连接
	db, err := mysql.GetDB()
	if err != nil {
		l.Logger.Errorf("数据库连接失败: %v", err)
		return nil, fmt.Errorf("数据库连接失败")
	}

	// 查询所有科室信息
	var depts []Department
	if err := db.Find(&depts).Error; err != nil {
		l.Logger.Errorf("查询科室列表失败: %v", err)
		return nil, err
	}

	// 构建响应数据
	resp := &patient.ListDepartmentsResponse{}
	for _, d := range depts {
		// 将数据库模型转换为proto响应格式
		resp.Departments = append(resp.Departments, &patient.Department{
			Id:          d.ID,
			Name:        d.Name,
			Description: d.Description,
		})
	}

	l.Logger.Infof("成功查询到 %d 个科室", len(resp.Departments))
	return resp, nil
}
