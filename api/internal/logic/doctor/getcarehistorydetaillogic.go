package doctor // doctor 逻辑包

import ( // 导入所需包
	"context" // 上下文包

	"github.com/qas491/hospital/api/internal/svc"   // 服务上下文
	"github.com/qas491/hospital/api/internal/types" // 类型定义

	"github.com/qas491/hospital/doctor_srv/doctor" // doctor rpc定义
	"github.com/zeromicro/go-zero/core/logx"       // 日志库
)

type GetCareHistoryDetailLogic struct { // 获取病例详情逻辑结构体
	logx.Logger                     // 日志记录器
	ctx         context.Context     // 请求上下文
	svcCtx      *svc.ServiceContext // 服务上下文
}

// NewGetCareHistoryDetailLogic 创建 GetCareHistoryDetailLogic 实例
func NewGetCareHistoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCareHistoryDetailLogic {
	return &GetCareHistoryDetailLogic{ // 返回结构体实例
		Logger: logx.WithContext(ctx), // 初始化日志
		ctx:    ctx,                   // 赋值上下文
		svcCtx: svcCtx,                // 赋值服务上下文
	}
}

// GetCareHistoryDetail 获取病例详情
func (l *GetCareHistoryDetailLogic) GetCareHistoryDetail(req *types.GetCareHistoryDetailReq) (*types.GetCareHistoryDetailResp, error) {
	rpcResp, err := l.svcCtx.DoctorRpc.GetCareHistoryDetail(l.ctx, &doctor.GetCareHistoryDetailReq{ // 调用RPC服务获取病例详情
		ChId: req.Ch_id, // 病例ID
	})
	if err != nil { // 如果有错误
		return nil, err // 返回错误
	}
	return &types.GetCareHistoryDetailResp{ // 返回响应结构体
		Code:    rpcResp.Code,                             // 状态码
		Message: rpcResp.Message,                          // 提示信息
		Detail:  convertToCareHistoryInfo(rpcResp.Detail), // 详情信息
	}, nil
}

// convertToCareHistoryInfo 转换RPC病例详情为本地类型
func convertToCareHistoryInfo(in *doctor.CareHistoryInfo) types.CareHistoryInfo {
	if in == nil { // 如果输入为空
		return types.CareHistoryInfo{} // 返回空结构体
	}
	return types.CareHistoryInfo{ // 构造本地病例详情
		Ch_id:         in.ChId,         // 病例ID
		User_id:       in.UserId,       // 用户ID
		User_name:     in.UserName,     // 用户名
		Patient_id:    in.PatientId,    // 患者ID
		Patient_name:  in.PatientName,  // 患者姓名
		Dept_id:       in.DeptId,       // 科室ID
		Dept_name:     in.DeptName,     // 科室名称
		Receive_type:  in.ReceiveType,  // 就诊类型
		Is_contagious: in.IsContagious, // 是否传染
		Care_time:     in.CareTime,     // 就诊时间
		Case_date:     in.CaseDate,     // 病例日期
		Reg_id:        in.RegId,        // 挂号ID
		Case_title:    in.CaseTitle,    // 病例标题
		Case_result:   in.CaseResult,   // 诊断结果
		Doctor_tips:   in.DoctorTips,   // 医生建议
		Remark:        in.Remark,       // 备注
	}
}
