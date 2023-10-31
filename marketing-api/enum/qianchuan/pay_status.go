package qianchuan

// PayStatus 支付属性
type PayStatus string

const (
	// PayStatus_PAID 已支付
	PayStatus_PAID PayStatus = "PAID"
	// PayStatus_UNPAID 未支付
	PayStatus_UNPAID PayStatus = "UNPAID"
)
