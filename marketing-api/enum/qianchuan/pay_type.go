package qianchuan

// PayType 订单类型
type PayType string

const (
	// PayType_DIRECT 直接转换订单
	PayType_DIRECT PayType = "DIRECT"
	// PayType_INDIRECT 间接转化订单
	PayType_INDIRECT PayType = "INDIRECT"
)
