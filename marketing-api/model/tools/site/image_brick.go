package site

// ImageBrick 图片组件
type ImageBrick struct {
	BaseBrick
	// ImageURL 图片链接，用户自行上传图片url，image_url 和图片ic_id 必须传一个。如果都传，取image_url值
	ImageURL string `json:"image_url,omitempty"`
	// IcID 图片ID(素材图片id)，对应【获取图片素材】接口获得的id; image_url和图片ic_id 必须传一个。如果都传，取image_url值。
	IcID string `json:"ic_id,omitempty"`
	// Tag 标签，用户自定义标注，类似图片字幕,默认为无
	Tag string `json:"tag,omitempty"`
	// BorderRadius 边框圆角，默认为0，范围：radius >= 0
	BorderRadius float64 `json:"border_radius,omitempty"`
	// Link 图片跳转链接信息
	Link *Link `json:"link,omitempty"`
	// IsCover 是否保持宽高比 ， 1: 表示保持宽高比， 0：不保持
	IsCover bool `json:"is_cover"`
}

// Type implement Brick interface
func (b ImageBrick) Type() BrickType {
	return XrPicture
}
