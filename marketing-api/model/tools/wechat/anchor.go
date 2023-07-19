package wechat

// Anchor 锚点组件信息
type Anchor struct {
	// IconImageURL 小游戏icon图片的url地址，图片尺寸60*60px,大小不超2MB
	IconImageURL string `json:"icon_image_url,omitempty"`
	// HeaderImageURL 顶部头图的url地址，大小不超过5MB的图片，推荐尺寸2:1
	HeaderImageURL string `json:"header_image_url,omitempty"`
	// Labels 小游戏标签，每个标签长度不超过6，最多支持6个标签
	Labels []string `json:"labels,omitempty"`
	// GuideText 引导文案，最大长度不超过14
	GuideText string `json:"guide_text,omitempty"`
	// ImagesVerticalURL 小游戏竖图的url地址，要求尺寸必须为3:5，否则报错。
	// 仅支持竖图和横图择其一类型上传，若同时上传两种类型的图片则报错
	// 小游戏横图或竖图上传最少3张，最多8张，大小不超过5MB
	ImagesVerticalURL []string `json:"images_vertical_url,omitempty"`
	// ImagesHorizontalURL 小游戏竖图的url地址，要求尺寸必须为3:5，否则报错。
	// 仅支持竖图和横图择其一类型上传，若同时上传两种类型的图片则报错
	// 小游戏横图或竖图上传最少3张，最多8张，大小不超过5MB
	ImagesHorizontalURL []string `json:"images_horizontal_url,omitempty"`
	// Introduction 小游戏简介，最大长度不超过45
	Introduction string `json:"introduction,omitempty"`
}
