package enum

// AudienceMode 人群定向模式
type AudienceMode string

const (
	// AudienceMode_AUTO 智能推荐
	AudienceMode_AUTO AudienceMode = "AUTO"
	// AudienceMode_CUSTOM 自定义
	AudienceMode_CUSTOM AudienceMode = "CUSTOM"
)
