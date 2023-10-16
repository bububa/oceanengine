package carousel

// Image 图片信息
type Image struct {
	// ImageID 图片id
	ImageID string `json:"image_id,omitempty"`
	// ImageMaterialID 图片素材id
	ImageMaterialID uint64 `json:"image_material_id,omitempty"`
	// Height 图片高度
	Height int `json:"height,omitempty"`
	// Width 图片宽度
	Width int `json:"width,omitempty"`
	// Ratio 图片比例
	Ratio float64 `json:"ratio,omitempty"`
	// WebURL 图片url
	WebURL string `json:"web_url,omitempty"`
	// URL 图片url
	URL string `json:"url,omitempty"`
}
