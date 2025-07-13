package doctor // doctor 逻辑包

import ( // 导入所需包
	"context" // 上下文包

	"github.com/qas491/hospital/api/internal/svc"   // 服务上下文
	"github.com/qas491/hospital/api/internal/types" // 类型定义

	"github.com/qas491/hospital/doctor_srv/doctor" // doctor rpc定义
	"github.com/zeromicro/go-zero/core/logx"       // 日志库
)

type GetCareHistoryListLogic struct { // 获取病例列表逻辑结构体
	logx.Logger                     // 日志记录器
	ctx         context.Context     // 请求上下文
	svcCtx      *svc.ServiceContext // 服务上下文
}

// NewGetCareHistoryListLogic 创建 GetCareHistoryListLogic 实例
func NewGetCareHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCareHistoryListLogic {
	return &GetCareHistoryListLogic{ // 返回结构体实例
		Logger: logx.WithContext(ctx), // 初始化日志
		ctx:    ctx,                   // 赋值上下文
		svcCtx: svcCtx,                // 赋值服务上下文
	}
}

// GetCareHistoryList 获取病例列表
func (l *GetCareHistoryListLogic) GetCareHistoryList(req *types.GetCareHistoryListReq) (*types.GetCareHistoryListResp, error) {
	rpcResp, err := l.svcCtx.DoctorRpc.GetCareHistoryList(l.ctx, &doctor.GetCareHistoryListReq{ // 调用RPC服务获取病例列表
		Page:        req.Page,         // 页码
		PageSize:    req.Page_size,    // 每页数量
		PatientId:   req.Patient_id,   // 患者ID
		PatientName: req.Patient_name, // 患者姓名
		UserId:      req.User_id,      // 用户ID
		DeptId:      req.Dept_id,      // 科室ID
		CaseDate:    req.Case_date,    // 病例日期
		StartTime:   req.Start_time,   // 开始时间
		EndTime:     req.End_time,     // 结束时间
	})
	if err != nil { // 如果有错误
		return nil, err // 返回错误
	}
	var list []types.CareHistoryInfo // 定义病例列表切片
	for _, c := range rpcResp.List { // 遍历RPC返回的列表
		list = append(list, convertToCareHistoryInfo(c)) // 转换并添加到本地列表
	}
	return &types.GetCareHistoryListResp{ // 返回响应结构体
		Code:    rpcResp.Code,    // 状态码
		Message: rpcResp.Message, // 提示信息
		List:    list,            // 病例列表
		Total:   rpcResp.Total,   // 总数
	}, nil
}
