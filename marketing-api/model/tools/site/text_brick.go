package site

// TextBrick 文本组件
type TextBrick struct {
	BaseBrick
	// Content 文本内容，长度至少为1
	Content string `json:"content,omitempty"`
	// Link 跳转链接信息
	Link *Link `json:"link,omitempty"`
}

// Type implement Brick interface
func (b TextBrick) Type() BrickType {
	return XrSimpleText
}
