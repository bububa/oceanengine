package enum

// AnchorRelatedType 原生锚点关联类型
type AnchorRelatedType string

const (
	// AnchorRelatedType_AUTO 自动生成
	AnchorRelatedType_AUTO AnchorRelatedType = "AUTO"
	// AnchorRelatedType_OFF 不启用
	AnchorRelatedType_OFF AnchorRelatedType = "OFF"
	// AnchorRelatedType_SELECT 手动选择
	AnchorRelatedType_SELECT AnchorRelatedType = "SELECT"
)
