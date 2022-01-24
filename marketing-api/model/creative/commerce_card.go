package creative

// CommerceCard 产品（商业卡）信息
type CommerceCard struct {
	// Title 产品卖点
	Title string `json:"title,omitempty"`
	// Source 产品名称
	Source string `json:"source,omitempty"`
	// ImageID 直播卡片图片信息。传入commerce_cards时必填。可通过调用【获取图片素材】获得。该图片信息作展示使用，可以自行选择您素材库中需要展示的图片进行上传。建议最佳：宽高比1:1, 224px*224px,小于300KB
	ImageID string `json:"image_id,omitempty"`
	// ImageInfo 素材信息
	ImageInfo *struct {
		// Width 宽度
		Width int `json:"width,omitempty"`
		// Height 高度
		Height int `json:"height,omitempty"`
		// WebUri 直播卡片图片信息
		WebUri string `json:"web_uri,omitempty"`
	} `json:"image_info,omitempty"`
}
