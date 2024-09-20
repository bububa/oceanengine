package site

// XrWechataBrick 微信小程序/微信小游戏组件
type WechatBrick struct {
	BaseBrick
	// InstanceID 小游戏ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// GamePath 小游戏路径参数
	GamePath string `json:"game_path,omitempty"`
	// Setting 设置
	Setting *WechatSetting `json:"setting,omitempty"`
}

type WechatSetting struct {
	// StyleType 允许值：
	// DESC_BUTTON 样式1（下图以微信小游戏组件为例）
	// WITHOUT_DESC_BUTTON样式2
	// BUTTON样式3
	StyleType int `json:"style_type,omitempty"`
	// Background style_type = DESC_BUTTON 样式1 时生效，样式-背景设置
	Background *WechatBackground `json:"background,omitempty"`
	// Image style_type = WITHOUT_DESC_BUTTON样式2 时生效，内容设置-背景图
	Image *WechatImage `json:"image,omitempty"`
	// IntroduceType 内容简介，
	// 允许值：text文本，label 标签
	IntroduceType string `json:"introduce_type,omitempty"`
	// Items 当introduce_type = label时，
	// 当label = 标签文案，最多支持传入2个
	Items []string `json:"items,omitempty"`
	// Title 小游戏/小程序名称
	Title *WechatTitle `json:"title,omitempty"`
	// Label 小游戏/小程序简介，当introduce_type = text时生效
	Label *WechatTitle `json:"label,omitempty"`
	// Avatar Logo 图片，宽高比建议为1:1，图片无水印
	Avatar *WechatImage `json:"avatar,omitempty"`
	// Button 按钮
	Button *WechatButton `json:"button,omitempty"`
}

type WechatBackground struct {
	// Type 背景设置
	// 允许值：color使用颜色；image使用图片
	Type string `json:"type,omitempty"`
	// Color 当type = color时传入生效
	Color string `json:"color,omitempty"`
	// Image 当type = image时传入生效
	Image *WechatImage `json:"image,omitempty"`
}

type WechatImage struct {
	// IcID 图片ID(素材图片id)，对应【获取图片素材】接口获得的image_url和图片ic_id 必须传一个。如果都传，取image_url值。
	IcID string `json:"ic_id,omitempty"`
	// ImageURL 图片链接，用户自行上传图片url，image_url 和图片ic_id 必须传一个。
	// 如果都传，取image_url值
	ImageURL string `json:"image_url,omitempty"`
}

type WechatTitle struct {
	// Color 小游戏/小程序名称颜色
	Color string `json:"color,omitempty"`
	// FontSize 小游戏/小程序名称字号
	FontSize float64 `json:"font_size,omitempty"`
	// Text 介绍文案
	Text string `json:"text,omitempty"`
}

type WechatButton struct {
	// Color 按钮文案颜色
	Color string `json:"color,omitempty"`
	// FontSize 按钮文案字号
	FontSize float64 `json:"font_size,omitempty"`
	// BackgroundColor 按钮填充颜色
	BackgroundColor string `json:"background_color,omitempty"`
	// BorderColor 按钮边框颜色
	BorderColor string `json:"border_color,omitempty"`
	// BorderWidth 按钮边框粗细
	BorderWidth float64 `json:"border_width,omitempty"`
	// BorderRadius 按钮圆角大小
	BorderRadius float64 `json:"border_radius,omitempty"`
}
