package enum

// AdRaiseOptType 一键起量操作类型
type AdRaiseOptType string

const (
	// AdRaiseOptType_CLICK_RAISE 启动一键起量
	AdRaiseOptType_CLICK_RAISE AdRaiseOptType = "CLICK_RAISE"
	// AdRaiseOptType_STOP_RAISE 关停一键起量
	AdRaiseOptType_STOP_RAISE AdRaiseOptType = "STOP_RAISE"
)
