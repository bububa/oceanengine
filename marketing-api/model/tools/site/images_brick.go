package site

// ImagesBrickGroupType 组图类型
type ImagesBrickGroupType string

const (
	// ImagesBrickGroupType_NORMAL 常规居中
	ImagesBrickGroupType_NORMAL ImagesBrickGroupType = "normal"
	// ImagesBrickGroupType_CAROUSEL 平铺拉伸
	ImagesBrickGroupType_CAROUSEL ImagesBrickGroupType = "carousel"
	// ImagesBrickGroupType_COVERFLOW 拼接覆盖
	ImagesBrickGroupType_COVERFLOW ImagesBrickGroupType = "coverflow"
)

// ImagesBrick 组图组件
type ImagesBrick struct {
	BaseBrick
	// GroupType 组图类型，默认normal;允许值：常规居中normal;平铺拉伸carousel;拼接覆盖coverflow
	GroupType ImagesBrickGroupType `json:"group_type,omitempty"`
	// AutoPlay 是否自动播放，默认自动播放;允许值：0（不自动），1（自动播放）
	AutoPlay int `json:"auto_play"`
	// GroupContent 组图内容列表
	GroupContent []ImageBrick `json:"group_content,omitempty"`
}

// Type implement Brick interface
func (b ImagesBrick) Type() BrickType {
	return XrPictureGroup
}
