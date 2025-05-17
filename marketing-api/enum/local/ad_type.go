package local

// AdType 广告类型
type AdType string

const (
	// AdType_ALL 不限
	AdType_ALL AdType = "ALL"
	// AdType_GENERAL 通投广告
	AdType_GENERAL AdType = "GENERAL"
	// AdType_SEARCHING 搜索广告
	AdType_SEARCHING AdType = "SEARCHING"
)
