package enum

type AdDeliveryRange string

const (
	AdDeliveryRange_DEFAULT   AdDeliveryRange = "DEFAULT"   // 默认
	AdDeliveryRange_UNION     AdDeliveryRange = "UNION"     // 只投放到资讯联盟（穿山甲）
	AdDeliveryRange_UNIVERSAL AdDeliveryRange = "UNIVERSAL" // 通投智选
)
