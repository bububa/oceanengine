package analyse

import "github.com/bububa/oceanengine/marketing-api/enum"

// Creative 商品创意数据
type Creative struct {
	// ImageMode 素材类型，可选值:
	// LARGE 大图
	// LARGE_VERTICAL 大图竖图
	// SMALL 小图
	// SQUARE 方图
	// UNION_SPLASH 穿山甲开屏图片
	// VIDEO_LARGE 横版视频
	// VIDEO_VERTICAL 竖版视频
	ImageMode enum.MaterialMode `json:"image_mode,omitempty"`
	// TitleMaterial 标题素材
	TitleMaterial *TitleMaterial `json:"title_material,omitempty"`
	// VideoMaterial 视频素材
	VideoMaterial *VideoMaterial `json:"video_material,omitempty"`
	Metrics
	// ProductInfo 商品信息
	ProductInfo *Product `json:"product_info,omitempty"`
}

// TitleMaterial 标题素材
type TitleMaterial struct {
	// Title 标题
	Title string `json:"title,omitempty"`
}

// VideoMaterial 视频素材
type VideoMaterial struct {
	// VideoID 视频id
	VideoID string `json:"video_id,omitempty"`
	// VideoDuration 视频时长
	VideoDuration float64 `json:"video_duration,omitempty"`
	// AwemeItemID 抖音主页视频id
	AwemeItemID uint64 `json:"aweme_item_id,omitempty"`
	// CoverImageHeight 视频封面高
	CoverImageHeight float64 `json:"cover_image_height,omitempty"`
	// CoverImageWidth 视频封面宽
	CoverImageWidth float64 `json:"cover_image_width,omitempty"`
	// CoverImageWebURL 视频封面url
	CoverImageWebURL string `json:"cover_image_web_url,omitempty"`
	// VideoURL 视频播放地址
	VideoURL []string `json:"video_url,omitempty"`
	// Source  视频来源，可选值:
	// AD 广告后台
	// AWEME_VIDEO 抖音原生视频
	Source string `json:"source,omitempty"`
}
