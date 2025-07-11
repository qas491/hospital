package logic

import (
	"context"
	"encoding/base64"
	"fmt"

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
	factory := until.NewAlipayFactory("9021000142690950", "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCQf++nIwcPt+buWMwZhZVlEa3qlY0igK+1JMlzFHXnOu27cFRnJNXWJxdJ9fdDTAaotkmHWb1DkWEBUmJkXCB1MaHxI+MMv5NqoZ105p8nUTlgodE5dTASie4cwDm9ptdRo3MGaxXWzeSOAcmAAiQ4wq+tB2+35cOzfsNNGnPxPpSxbDFCe9cYW3Rm57utu6zv5g29sblU/aD4cPuM+ZLvcnXGUZZ+B2XuwV4n3HMVQI43A/kly8H1UuPh63OBoCvlZ/K2vvvVnySkm9hdxXFv/oiL4SY6wIHTAQGxyscZLqYSATsumc882TJ15kJmtg77MWy7paq6ZA+0EexSxhgNAgMBAAECggEAW0kXYyUTvvBU6VobhjwHxxPGJR5ZTOSzG+KjiRnx6iQmL3WlPIvesy163vSOQmtzAX43MVSV5mufNrCPDAvPTXoPbkFXnKQiQwjaahGPlc0QTGWtwXiw5+VPEca2M4OFH0P81J8t6sejjbq/SykPLPSA+vRptWlnmquIQdtmR45oZFPBb4r5EmuJ6J7X1lYr/R/jga0Wa5mS80HJ0i882pAEL4f35cQgJujm+frk5aBdjIwZ+H6HtcAdFcRUQdGsA00x3JV20J8I+cRt+EbQoFpYmjbzCOrkPtxa/yG5gOIUJtQ1buv+aawv92e63Q+VdtVZxYc59zQ6GLgEbheZwQKBgQDDEe6baqlDgJyXi6cKqeUuz2SsZTUthan44H/7ilU1fNL4uLadTNS4tCzebjFpGT2qc/EansONMeb8OHecsTz21d85rwrGCBpWfNtV6zvgwNkuuB9kl1sUP+pqmd7mKxXoqiVUTNg7HnXoi41QQR0fDMA7nvmJT1vPQkk3sz3D+QKBgQC9olV+ktO5CtytK4geKtEyoi6LkVggRB3Vf5S200JdtWv2/Xp6RqEJSpnBYt47Kbj+Bd4cPsiXX11kz+qFUUjXwupyzUOHpOVEUzlel6gqs81VDNsjH+cyB/kiS2YeOjDMfL5uRrhhyfUqtGs1qdFc0PrgloOL8N/OGPRQto8RtQKBgFrs/N+MtCE2zccF3XLnBmDvYunIsyTo7PWJD57cOCOab2xoDRb9PRprQY7cpkNq9IeDS5sw7c3euOBQIdDz/IrB3i8xHEB6fmxZ/pLp9xsoSROx11A4DMg/krDl0DRRWQD+bjABMbk81ZDzm4cREtXqR6CC1aXfS9gr+Zzw+5VhAoGBALKLDQPHKiCm4W8J9XaxyZVqrXgquFZoy84f+NzJu0qPbb92mMJXjXc6DdnORH5fohVJYP4m/qXG3B/2wlATYAoFMsg0CsjDsDjMQs+U6niKIWFaYViIyRyJ9T8czmDXWOqu6HzbiO5JP9OdWvODl+NTv2GFVQWFHLLnO+BiggGhAoGAbhx5fMhNq3u5qH2H61K8VIzPKHuBf4V1vgjzb40OeN6k5/hUswEVZy3LFhFZ3/9pdJXVspOF6+QCClBPmuVNvx1a311IspfetJhh/EeExQAyFk2l94MHJKfs633Hdi4Tw9cUFNs8SXaBtkB9VSI0jqZdBOviNy4nq4w/dcu2hBI=", false)
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
