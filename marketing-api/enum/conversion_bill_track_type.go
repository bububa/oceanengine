package enum

// ConversionBillTrackType 转化支付方式
type ConversionBillTrackType = int

const (
	// ConversionBillTrackType_WECHAT 微信支付
	ConversionBillTrackType_WECHAT ConversionBillTrackType = 1
	// ConversionBillTrackType_ALIPAY 支付宝支付
	ConversionBillTrackType_ALIPAY ConversionBillTrackType = 2
)
