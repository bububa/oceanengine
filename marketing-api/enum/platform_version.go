package enum

// PlatformVersion 巨量引擎平台版本
type PlatformVersion string

const (
	// PlatformVersion_V1 原版
	PlatformVersion_V1 PlatformVersion = "V1"
	// PlatformVersion_V2 巨量引擎体验版
	PlatformVersion_V2 PlatformVersion = "V2"
	// PlatformVersion_ALL 原版+巨量引擎体验版（默认值）
	PlatformVersion_ALL PlatformVersion = "ALL"
)
