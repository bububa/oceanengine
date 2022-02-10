package enum

// CreativeDisplayMode 创意展现方式
type CreativeDisplayMode string

const (
	// CREATIVE_DISPLAY_MODE_CTR 优选(优先投放预估点击率高的创意素材)
	CREATIVE_DISPLAY_MODE_CTR CreativeDisplayMode = "CREATIVE_DISPLAY_MODE_CTR"
	// CREATIVE_DISPLAY_MODE_RANDOM 轮播
	CREATIVE_DISPLAY_MODE_RANDOM CreativeDisplayMode = "CREATIVE_DISPLAY_MODE_RANDOM"
)
