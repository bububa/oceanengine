package enum

// ExternalActionSource 转化目标创建来源
type ExternalActionSource string

const (
	// ExternalActionSource_SOURCE_TETRIS 建站 - 预定义
	ExternalActionSource_SOURCE_TETRIS ExternalActionSource = "SOURCE_TETRIS"
	// ExternalActionSource_SOURCE_METEOR 转化跟踪 - 自定义
	ExternalActionSource_SOURCE_METEOR ExternalActionSource = "SOURCE_METEOR"
	// ExternalActionSource_SOURCE_CONFIG 配置固定返回
	ExternalActionSource_SOURCE_CONFIG ExternalActionSource = "SOURCE_CONFIG"
)
