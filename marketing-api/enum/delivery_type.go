package enum

// DeliveryType 投放类型
type DeliveryType string

const (
	// DeliveryType_NORMAL 常规投放
	DeliveryType_NORMAL DeliveryType = "NORMAL"
	// DeliveryType_DURATION 周期稳投（目前仅支持搜索广告）
	DeliveryType_DURATION DeliveryType = "DURATION"
	// DeliveryType_UBX_INTELLIGENT UBA智能托管
	DeliveryType_UBX_INTELLIGENT DeliveryType = "UBX_INTELLIGENT"
)
