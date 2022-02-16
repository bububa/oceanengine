package coupon

// ActivityType 卡券类型
type ActivityType string

const (
	// ActivityType_DIRECT_NEED_PHONE 直接发券，收集手机号，need_phone须为true
	ActivityType_DIRECT_NEED_PHONE ActivityType = "DIRECT_NEED_PHONE"
	// ActivityType_DIRECT_NOT_NEED_PHONE 直接发券，不收集手机号，need_phone须为false 并且 BindFormId会被忽略
	ActivityType_DIRECT_NOT_NEED_PHONE ActivityType = "DIRECT_NOT_NEED_PHONE"
)
