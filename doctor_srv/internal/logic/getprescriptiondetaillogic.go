package logic

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/qas491/hospital/doctor_srv/configs"

	"github.com/qas491/hospital/doctor_srv/until"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/qas491/hospital/doctor_srv/internal/svc"
	"github.com/qas491/hospital/doctor_srv/model/mysql"
	"github.com/skip2/go-qrcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrescriptionDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrescriptionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrescriptionDetailLogic {
	return &GetPrescriptionDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取处方详情
func (l *GetPrescriptionDetailLogic) GetPrescriptionDetail(in *doctor.GetPrescriptionDetailReq) (*doctor.GetPrescriptionDetailResp, error) {
	var order mysql.HisCareOrder
	err := l.svcCtx.DB.Where("co_id = ?", in.CoId).First(&order).Error
	if err != nil {
		return &doctor.GetPrescriptionDetailResp{
			Code:    1,
			Message: "未找到处方: " + err.Error(),
		}, nil
	}
	var items []mysql.HisCareOrderItem
	l.svcCtx.DB.Where("co_id = ?", in.CoId).Find(&items)
	protoItems := make([]*doctor.PrescriptionItem, 0, len(items))
	for _, it := range items {
		protoItems = append(protoItems, &doctor.PrescriptionItem{
			ItemId:    it.ItemID,
			ItemRefId: it.ItemRefID,
			ItemName:  it.ItemName,
			ItemType:  it.ItemType,
			Num:       it.Num,
			Price:     it.Price,
			Amount:    it.Amount,
			Remark:    it.Remark,
			Status:    it.Status,
		})
	}

	// 生成支付宝支付URL
	factory := until.NewAlipayFactory(configs.WiseConfig.Alipay.AppId, configs.WiseConfig.Alipay.Key, false)
	payURL, err := factory.Pay(
		"处方支付",                        // subject
		order.CoID,                    // outTradeNo
		FormatAmount(order.AllAmount), // totalAmount
		"http://your_notify_url",      // notifyURL
		"http://your_return_url",      // returnURL
	)
	if err != nil {
		return &doctor.GetPrescriptionDetailResp{
			Code:    2,
			Message: "支付宝支付URL生成失败: " + err.Error(),
		}, nil
	}

	// 生成二维码
	png, err := qrcode.Encode(payURL, qrcode.Medium, 256)
	if err != nil {
		return &doctor.GetPrescriptionDetailResp{
			Code:    3,
			Message: "二维码生成失败: " + err.Error(),
		}, nil
	}
	qrBase64 := base64.StdEncoding.EncodeToString(png)

	resp := &doctor.GetPrescriptionDetailResp{
		Code:    0,
		Message: "ok",
		Detail: &doctor.PrescriptionInfo{
			CoId:        order.CoID,
			CoType:      order.CoType,
			UserId:      order.UserID,
			PatientId:   order.PatientID,
			PatientName: order.PatientName,
			ChId:        order.ChID,
			AllAmount:   order.AllAmount,
			CreateBy:    order.CreateBy,
			CreateTime:  FormatTime(order.CreateTime),
			UpdateBy:    order.UpdateBy,
			UpdateTime:  FormatTime(order.UpdateTime),
			Items:       protoItems,
		},
		Qrcode: qrBase64,
	}
	return resp, nil
}

// FormatAmount 格式化金额为字符串
func FormatAmount(amount float64) string {
	return fmt.Sprintf("%.2f", amount)
}
