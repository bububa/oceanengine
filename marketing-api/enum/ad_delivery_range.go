package enum

// AdDeliveryRange 广告投放范围
type AdDeliveryRange string

const (
	// AdDeliveryRange_DEFAULT 默认
	AdDeliveryRange_DEFAULT AdDeliveryRange = "DEFAULT"
	// AdDeliveryRange_UNION 只投放到资讯联盟（穿山甲）
	AdDeliveryRange_UNION AdDeliveryRange = "UNION"
	// AdDeliveryRange_UNIVERSAL 通投智选
	AdDeliveryRange_UNIVERSAL AdDeliveryRange = "UNIVERSAL"
)
