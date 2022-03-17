package enum

// LaunchTargetType 投放类型
type LaunchTargetType string

const (
	// LaunchTargetType_LIVE_CONVERT 直播间转化
	LaunchTargetType_LIVE_CONVERT LaunchTargetType = "LIVE_CONVERT"
	// LaunchTargetType_APP 应用下载
	LaunchTargetType_APP LaunchTargetType = "APP"
	// LaunchTargetType_EXTERNAL 线索收集
	LaunchTargetType_EXTERNAL LaunchTargetType = "EXTERNAL"
)
