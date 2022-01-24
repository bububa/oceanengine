package enum

// DpaVideoTemplateType 商品库视频生成类型
type DpaVideoTemplateType string

const (
	// DPA_VIDEO_TEMPLATE_SMART 优选商品库视频（自动根据商品库商品图片生成）
	DPA_VIDEO_TEMPLATE_SMART DpaVideoTemplateType = "DPA_VIDEO_TEMPLATE_SMART"
	// DPA_VIDEO_TEMPLATE_CUSTOM 自定义商品库视频
	DPA_VIDEO_TEMPLATE_CUSTOM DpaVideoTemplateType = "DPA_VIDEO_TEMPLATE_CUSTOM"
)
