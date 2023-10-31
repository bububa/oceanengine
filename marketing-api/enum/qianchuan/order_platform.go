package qianchuan

// OrderPlatform 下单平台
type OrderPlatform string

const (
	// OrderPlatform_ALL 全部
	OrderPlatform_ALL OrderPlatform = "ALL"
	// OrderPlatform_QIANCHUAN  千川pc（默认）
	OrderPlatform_QIANCHUAN OrderPlatform = "QIANCHUAN"
	// OrderPlatform_ECP_AWEME 小店随心推
	OrderPlatform_ECP_AWEME OrderPlatform = "ECP_AWEME"
)
