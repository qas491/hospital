package until

import (
	"github.com/smartwalle/alipay/v3"
)

// Payment 支付接口
// 可扩展微信、银联等
type Payment interface {
	Pay(subject, outTradeNo, totalAmount, notifyURL, returnURL string) (string, error)
}

// AlipayFactory 支付宝工厂
type AlipayFactory struct {
	appId  string
	priKey string
	isProd bool
}

func NewAlipayFactory(appId, priKey string, isProd bool) *AlipayFactory {
	return &AlipayFactory{
		appId:  appId,
		priKey: priKey,
		isProd: isProd,
	}
}

// Alipay 支付宝支付实现
func (a *AlipayFactory) Pay(subject, outTradeNo, totalAmount, notifyURL, returnURL string) (string, error) {
	client, err := alipay.New(a.appId, a.priKey, a.isProd)
	if err != nil {
		return "", err
	}
	var p = alipay.TradeWapPay{}
	p.NotifyURL = notifyURL
	p.ReturnURL = returnURL
	p.Subject = subject
	p.OutTradeNo = outTradeNo
	p.TotalAmount = totalAmount
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(p)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}

// Example usage:
// func main() {
// 	factory := NewAlipayFactory("9021000142690950", "<privateKey>", false)
// 	payURL, err := factory.Pay("标题", "唯一单号", "10.00", "http://notify", "http://return")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(payURL)
// }
