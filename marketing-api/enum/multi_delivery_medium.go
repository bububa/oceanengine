package enum

// MultiDeliveryMedium 多投放载体，仅当landing_type = LINK 销售线索推广目的下会返回
type MultiDeliveryMedium string

const (
	// MultiDeliveryMedium_ORANGE_AND_AWEME 优选投放橙子落地页和抖音主页
	MultiDeliveryMedium_ORANGE_AND_AWEME MultiDeliveryMedium = "ORANGE_AND_AWEME"
)
