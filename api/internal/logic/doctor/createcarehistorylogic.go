// CreateCareHistoryLogic 处理创建病例的业务逻辑
package doctor // doctor 逻辑包

import ( // 导入所需包
	"context" // 上下文包

	"github.com/qas491/hospital/api/internal/svc"   // 服务上下文
	"github.com/qas491/hospital/api/internal/types" // 类型定义

	"github.com/qas491/hospital/doctor_srv/doctor" // doctor rpc定义
	"github.com/zeromicro/go-zero/core/logx"       // 日志库
)

type CreateCareHistoryLogic struct { // 创建病例逻辑结构体
	logx.Logger                     // 日志记录器
	ctx         context.Context     // 请求上下文
	svcCtx      *svc.ServiceContext // 服务上下文
}

// NewCreateCareHistoryLogic 创建 CreateCareHistoryLogic 实例
func NewCreateCareHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCareHistoryLogic {
	return &CreateCareHistoryLogic{ // 返回结构体实例
		Logger: logx.WithContext(ctx), // 初始化日志
		ctx:    ctx,                   // 赋值上下文
		svcCtx: svcCtx,                // 赋值服务上下文
	}
}

// CreateCareHistory 创建病例
func (l *CreateCareHistoryLogic) CreateCareHistory(req *types.CreateCareHistoryReq) (*types.CreateCareHistoryResp, error) {
	rpcResp, err := l.svcCtx.DoctorRpc.CreateCareHistory(l.ctx, &doctor.CreateCareHistoryReq{ // 调用RPC服务创建病例
		ChId:         req.Ch_id,         // 病例ID
		UserId:       req.User_id,       // 用户ID
		UserName:     req.User_name,     // 用户名
		PatientId:    req.Patient_id,    // 患者ID
		PatientName:  req.Patient_name,  // 患者姓名
		DeptId:       req.Dept_id,       // 科室ID
		DeptName:     req.Dept_name,     // 科室名称
		ReceiveType:  req.Receive_type,  // 就诊类型
		IsContagious: req.Is_contagious, // 是否传染
		CaseDate:     req.Case_date,     // 病例日期
		RegId:        req.Reg_id,        // 挂号ID
		CaseTitle:    req.Case_title,    // 病例标题
		CaseResult:   req.Case_result,   // 诊断结果
		DoctorTips:   req.Doctor_tips,   // 医生建议
		Remark:       req.Remark,        // 备注
	})
	if err != nil { // 如果有错误
		return nil, err // 返回错误
	}
	return &types.CreateCareHistoryResp{ // 返回响应结构体
		Code:    rpcResp.Code,    // 状态码
		Message: rpcResp.Message, // 提示信息
		Ch_id:   rpcResp.ChId,    // 病例ID
	}, nil
}
