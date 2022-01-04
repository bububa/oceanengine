package enum

// AdRaiseStatus 一键起量状态
type AdRaiseStatus string

const (
	// AdRaiseStatus_UNKNOWN 未知状态
	AdRaiseStatus_UNKNOWN AdRaiseStatus = ""
	// AdRaiseStatus_ENABLE_RAISE 可以使用一键起量
	AdRaiseStatus_ENABLE_RAISE AdRaiseStatus = "ENABLE_RAISE"
	// AdRaiseStatus_DISABLE_RAISE 不可以使用一键起量
	AdRaiseStatus_DISABLE_RAISE AdRaiseStatus = "DISABLE_RAISE"
	// AdRaiseStatus_RAISING 一键起量中
	AdRaiseStatus_RAISING AdRaiseStatus = "RAISING"
	// AdRaiseStatus_FINISH_RAISE 一键起量完成
	AdRaiseStatus_FINISH_RAISE AdRaiseStatus = "FINISH_RAISE"
	// AdRaiseStatus_ENABLE_PRERAISE 可预约
	AdRaiseStatus_ENABLE_PRERAISE AdRaiseStatus = "ENABLE_PRERAISE"
	// AdRaiseStatus_HAS_PRERAISE 已预约
	AdRaiseStatus_HAS_PRERAISE AdRaiseStatus = "HAS_PRERAISE"
)
