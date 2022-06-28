package dpa

import "github.com/bububa/oceanengine/marketing-api/enum"

// Video 视频模板
type Video struct {
	// ProductPlatformID 商品库id
	ProductPlatformID uint64 `json:"product_platform_id,omitempty"`
	// ProductID 商品库商品id
	ProductID uint64 `json:"product_id,omitempty"`
	// PackageID 商品库视频模板id
	PackageID string `json:"package_id,omitempty"`
	// Name 商品库视频模板名称
	Name string `json:"name,omitempty"`
	// CreateTime 创建时间，格式 %Y-%m-%d %H:%M:%S
	CreateTime string `json:"create_time,omitempty"`
	// ImageInfo 视频模板封面图信息
	ImageInfo *ImageInfo `json:"image_info,omitempty"`
	// VideoInfo 商品库视频模板视频信息
	VideoInfo *VideoInfo `json:"video_info,omitempty"`
}

// ImageInfo 视频模板封面图信息
type ImageInfo struct {
	// Height 图片高度
	Height int `json:"height,omitempty"`
	// ImageID 图片id
	ImageID string `json:"image_id,omitempty"`
	// Width 图片宽度
	Width int `json:"width,omitempty"`
}

// VideoInfo 商品库视频模板视频信息
type VideoInfo struct {
	// Height 视频高度
	Height int `json:"height,omitempty"`
	// Width 视频宽度
	Width int `json:"width,omitempty"`
	// ImageMode 视频类型，详见【附录-枚举值-素材类型】
	// 可选值: CREATIVE_IMAGE_MODE_VIDEO：横版视频,CREATIVE_IMAGE_MODE_VIDEO_VERTICAL：竖版视频
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// Status 视频状态，可选值: AVAILABLE：可用,UNAVAILABLE：不可用
	Status string `json:"status,omitempty"`
	// VideoID 视频id
	VideoID string `json:"video_id,omitempty"`
}
