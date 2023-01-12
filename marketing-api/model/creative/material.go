package creative

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// Material 素材信息
type Material struct {
	// ID 创意ID
	ID model.Uint64 `json:"id,omitempty"`
	// AdID 广告计划ID
	AdID model.Uint64 `json:"ad_id,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID model.Uint64 `json:"advertiser_id,omitempty"`
	// Title 创意素材标题
	Title string `json:"title,omitempty"`
	// Status 创意状态
	Status enum.CreativeStatus `json:"status,omitempty"`
	// OptStatus 创意操作状态
	OptStatus enum.CreativeOptStatus `json:"op_status,omitempty"`
	// ImageMode 创意类型
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// ImageInfo 图片素材信息
	ImageInfo []struct {
		// Url 图片链接
		Url string `json:"url,omitempty"`
		// Width 图片宽度
		Width int `json:"width,omitempty"`
		// Height 图片高度
		Height int `json:"height,omitempty"`
	} `json:"image_info,omitempty"`
	// ImageID 视频素材封面
	ImageID string `json:"image_id,omitempty"`
	// VideoID 视频ID
	VideoID string `json:"video_id,omitempty"`
}
