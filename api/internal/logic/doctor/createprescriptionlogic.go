// CreatePrescriptionLogic 处理创建处方的业务逻辑
package doctor // doctor 逻辑包

import ( // 导入所需包
	"context" // 上下文包

	"github.com/qas491/hospital/api/internal/svc"   // 服务上下文
	"github.com/qas491/hospital/api/internal/types" // 类型定义
	"github.com/qas491/hospital/doctor_srv/doctor"  // doctor rpc定义

	"github.com/zeromicro/go-zero/core/logx" // 日志库
)

type CreatePrescriptionLogic struct { // 创建处方逻辑结构体
	logx.Logger                     // 日志记录器
	ctx         context.Context     // 请求上下文
	svcCtx      *svc.ServiceContext // 服务上下文
}

// NewCreatePrescriptionLogic 创建 CreatePrescriptionLogic 实例
func NewCreatePrescriptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePrescriptionLogic {
	return &CreatePrescriptionLogic{ // 返回结构体实例
		Logger: logx.WithContext(ctx), // 初始化日志
		ctx:    ctx,                   // 赋值上下文
		svcCtx: svcCtx,                // 赋值服务上下文
	}
}

// CreatePrescription 创建处方
func (l *CreatePrescriptionLogic) CreatePrescription(req *types.CreatePrescriptionReq) (*types.CreatePrescriptionResp, error) {
	rpcReq := &doctor.CreatePrescriptionReq{ // 构造RPC请求体
		CoId:        req.Co_id,                                   // 处方ID
		CoType:      req.Co_type,                                 // 处方类型
		UserId:      req.User_id,                                 // 用户ID
		PatientId:   req.Patient_id,                              // 患者ID
		PatientName: req.Patient_name,                            // 患者姓名
		ChId:        req.Ch_id,                                   // 病例ID
		AllAmount:   req.All_amount,                              // 总金额
		CreateBy:    req.Create_by,                               // 创建人
		Items:       convertToDoctorPrescriptionItems(req.Items), // 处方明细
	}
	rpcResp, err := l.svcCtx.DoctorRpc.CreatePrescription(l.ctx, rpcReq) // 调用RPC服务创建处方
	if err != nil {                                                      // 如果有错误
		return nil, err // 返回错误
	}
	return &types.CreatePrescriptionResp{ // 返回响应结构体
		Code:    rpcResp.Code,    // 状态码
		Message: rpcResp.Message, // 提示信息
		Co_id:   rpcResp.CoId,    // 处方ID
	}, nil
}

// convertToDoctorPrescriptionItems 转换处方明细为RPC格式
func convertToDoctorPrescriptionItems(items []types.PrescriptionItem) []*doctor.PrescriptionItem {
	var result []*doctor.PrescriptionItem // 定义结果切片
	for _, item := range items {          // 遍历每个明细项
		result = append(result, &doctor.PrescriptionItem{ // 构造明细项
			ItemId:    item.Item_id,     // 明细ID
			ItemRefId: item.Item_ref_id, // 关联ID
			ItemName:  item.Item_name,   // 明细名称
			ItemType:  item.Item_type,   // 明细类型
			Num:       item.Num,         // 数量
			Price:     item.Price,       // 单价
			Amount:    item.Amount,      // 金额
			Remark:    item.Remark,      // 备注
			Status:    item.Status,      // 状态
		})
	}
	return result // 返回结果切片
}
