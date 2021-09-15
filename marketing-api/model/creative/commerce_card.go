package creative

// CommerceCard 产品（商业卡）信息
type CommerceCard struct {
	// Title 产品卖点
	Title string `json:"title,omitempty"`
	// Source 产品名称
	Source string `json:"source,omitempty"`
	// ImageInfo 素材信息
	ImageInfo *struct {
		// Width 宽度
		Width int `json:'width,omitempty'`
		// Height 高度
		Height int `json:"height,omitempty"`
		// WebUri 直播卡片图片信息
		WebUri string `json:'web_uri,omitempty'`
	} `json:"image_info,omitempty"`
}
