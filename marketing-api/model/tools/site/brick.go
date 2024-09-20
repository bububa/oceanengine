package site

type BrickType string

const (
	// XrVideo 视频组件
	XrVideo BrickType = "XrVideo"
	// XrPicture 图片组件
	XrPicture BrickType = "XrPicture"
	// XrPictureGroup 组图组件
	XrPictureGroup BrickType = "XrPictureGroup"
	// XrSimpleText 文本组件
	XrSimpleText BrickType = "XrSimpleText"
	// XrRichText 富文本组件
	XrRichText BrickType = "XrRichText"
	// XrForm 表单组件
	XrForm BrickType = "XrForm"
	// XrButton 按钮组件
	XrButton BrickType = "XrButton"
	// XrWechataApplet 微信小游戏组件
	XrWechataApplet BrickType = "XrWechataApplet"
	// XrWechatGame 微信小程序组件类型
	XrWechatGame BrickType = "XrWechatGame"
)

// Brick 站点业务数据
type Brick interface {
	Type() BrickType
}

// BaseBrick 基础组件信息
type BaseBrick struct {
	// Name 组件类型
	Name string `json:"name,omitempty"`
	// X 横轴，代表当前组件在布局中的左右位置，默认为0
	X int `json:"x,omitempty"`
	// Y 纵轴，代表当前组件在布局中的上下位置，默认为0
	Y int `json:"y,omitempty"`
	// Width 宽度，默认为375，范围：width >= 1
	Width int `json:"width,omitempty"`
	// Height 高度，默认为225，范围：height >= 1
	Height int `json:"height,omitempty"`
	// Float 站点组件浮动方式: none（默认为无);其他值可选:top（置顶）;right（居右）;bottom（置底）
	Float string `json:"float,omitempty"`
}

func (b BaseBrick) Type() BrickType {
	return BrickType(b.Name)
}
