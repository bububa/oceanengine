package enum

// LaunchType 调起方式
type LaunchType string

const (
	// LaunchType_DIRECT_OPEN 直接调起
	LaunchType_DIRECT_OPEN LaunchType = "DIRECT_OPEN"
	// LaunchType_EXTERNAL_OPEN 落地页调起
	LaunchType_EXTERNAL_OPEN LaunchType = "EXTERNAL_OPEN"
)
