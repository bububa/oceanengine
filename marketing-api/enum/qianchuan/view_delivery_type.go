package qianchuan

// ViewDeliveryType 资金池类型，允许值：
type ViewDeliveryType string

const (
	// ViewDeliveryType_ALL 全部（默认值）
	ViewDeliveryType_ALL ViewDeliveryType = "ALL"
	// ViewDeliveryType_DEFAULT 通用
	ViewDeliveryType_DEFAULT ViewDeliveryType = "DEFAULT"
	// ViewDeliveryType_BRAND 品牌
	ViewDeliveryType_BRAND ViewDeliveryType = "BRAND"
)
