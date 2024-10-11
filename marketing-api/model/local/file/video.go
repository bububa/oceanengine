package file

import "github.com/bububa/oceanengine/marketing-api/enum/local"

// Video 视频信息
type Video struct {
	// VideoID 视频id
	VideoID string `json:"video_id,omitempty"`
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// Size 视频大小
	Size int64 `json:"size,omitempty"`
	// VideoSiganture 视频md5
	VideoSiganture string `json:"video_siganture,omitempty"`
	// Width 视频宽
	Width int64 `json:"width,omitempty"`
	// Height 视频高
	Height int64 `json:"height,omitempty"`
	// VideoURL 视频地址
	VideoURL string `json:"video_url,omitempty"`
	// Duration 视频时长
	Duration float64 `json:"duration,omitempty"`
	// VideoName 视频名称
	VideoName string `json:"video_name,omitempty"`
	// PosterURL 视频首帧截图
	PosterURL string `json:"poster_url,omitempty"`
	// MaterialProperties 素材标签，枚举值：
	// COPY 搬运风险
	// FIRST_PUBLISH 首发
	// HIGH_QUALITY 优质
	// LOW_QUALITY 低质
	// SIMILAR 同质化风险
	MaterialProperties []string `json:"material_properties,omitempty"`
	// ImageMode 视频类型，枚举值：
	// IMAGE_MODE_VIDEO 横版视频
	// IMAGE_MODE_VIDEO_VERTICAL 竖版视频
	ImageMode local.ImageMode `json:"image_mode,omitempty"`
	// Source  视频来源，枚举值：
	// BP_PLATFORM 巨量引擎工作平台共享视频
	// CREATIVE_AIGC 即创
	// LOCAL_ADS_UPLOAD 本地上传
	// STAR 星图平台
	// MAPI MAPI接口上传
	Source local.MaterialSource `json:"source,omitempty"`
	// CreateTime 素材的上传时间，格式：yyyy-mm-dd HH:mm:ss
	CreateTime string `json:"create_time,omitempty"`
}

// VideoAweme 抖音视频
type VideoAweme struct {
	// ItemID 抖音视频ID
	ItemID uint64 `json:"item_id,omitempty"`
	// Title 视频标题
	Title string `json:"title,omitempty"`
	// VideoID 视频ID
	VideoID string `json:"video_id,omitempty"`
	// AwemeID 抖音号ID
	AwemeID string `json:"aweme_id,omitempty"`
	// ImageMode 视频类型，枚举值：
	// IMAGE_MODE_VIDEO 横版视频
	// IMAGE_MODE_VIDEO_VERTICAL 竖版视频
	ImageMode local.ImageMode `json:"image_mode,omitempty"`
	// CoverImageURL 视频封面图片地址
	CoverImageURL string `json:"cover_image_url,omitempty"`
	// AwemeVideoURL 视频播放地址
	AwemeVideoURL string `json:"aweme_video_url,omitempty"`
	// NotDeliveryReason 不可投放原因
	NotDeliveryReason string `json:"not_delivery_reason,omitempty"`
	// CanDelivery 视频是否可投放
	// true 可投放
	// false 不可投放
	CanDelivery bool `json:"can_delivery,omitempty"`
	// LegoMaterialID 素材id
	LegoMaterialID uint64 `json:"lego_material_id,omitempty"`
	// VideoWidth 视频宽度
	VideoWidth int64 `json:"video_width,omitempty"`
	// VideoHeight 视频高度
	VideoHeight int64 `json:"video_height,omitempty"`
	// Duration 视频时长
	Duration string `json:"duration,omitempty"`
}
