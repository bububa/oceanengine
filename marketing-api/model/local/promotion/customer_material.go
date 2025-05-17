package promotion

import "github.com/bububa/oceanengine/marketing-api/enum/local"

// CustomerMaterial 视频素材列表
type CustomerMaterial struct {
	// ImageMode 素材类型 可选值:
	// IMAGE_MODE_VIDEO 横版视频
	// IMAGE_MODE_VIDEO_VERTICAL 竖版视频
	ImageMode local.ImageMode `json:"image_mode,omitempty"`
	// TitleMaterial 标题素材
	// 仅针对素材库视频生效且必填，即传入video_id时，该字段生效且必填
	// 当传入item_id时，以抖音主页视频的标题为主，该字段不生效
	TitleMaterial *TitleMaterial `json:"title_material,omitempty"`
	// VideoMaterial 视频素材
	VideoMaterial *VideoMaterial `json:"video_material,omitempty"`
}

// TitleMaterial 标题素材
type TitleMaterial struct {
	// Title 标题，标题长度要求5-55个字（两个英文字符占1个字）
	Title string `json:"title,omitempty"`
	// LegoMaterialID 标题素材库id
	LegoMaterialID uint64 `json:"lego_material_id,omitempty"`
	// MaterialID 标题素材id
	MaterialID uint64 `json:"material_id,omitempty"`
}

// VideoMaterial 视频素材
type VideoMaterial struct {
	// VideoID 视频ID
	VideoID string `json:"video_id,omitempty"`
	// CoverWebURL 封面图片URI
	CoverWebURL string `json:"cover_web_url,omitempty"`
	// AwemeItemID 抖音短视频ID，推广抖音主页视频时传入
	// 注意：如果视频为抖音主页视频，则该字段必填
	// 如果和video_id同时传，以aweme_item_id为准
	AwemeItemID uint64 `json:"aweme_item_id,omitempty"`
	// LegoMaterialID 视频素材库id
	LegoMaterialID uint64 `json:"lego_material_id,omitempty"`
	// MaterialID 视频素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// ImageMode 素材类型，枚举值：
	// IMAGE_MODE_VIDEO 横版视频
	// IMAGE_MODE_VIDEO_VERTICAL 竖版视频
	ImageMode local.ImageMode `json:"image_mode,omitempty"`
	// VideoDuration 视频长度
	VideoDuration int64 `json:"video_duration,omitempty"`
	// VideoHeight 视频高度
	VideoHeight int64 `json:"video_height,omitempty"`
	// VideoWidth 视频宽度
	VideoWidth int64 `json:"video_width,omitempty"`
	// VideoPlayURL 视频播放链接
	VideoPlayURL string `json:"video_play_url,omitempty"`
	// CoverImageHeight 封面图片高度
	CoverImageHeight int64 `json:"cover_image_height,omitempty"`
	// CoverImageWidth 封面图片宽度
	CoverImageWidth int64 `json:"cover_image_width,omitempty"`
	// CoverWebURI 封面图片uri
	CoverWebURI string `json:"cover_web_uri,omitempty"`
}
