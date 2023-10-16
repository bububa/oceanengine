package creative

import "github.com/bububa/oceanengine/marketing-api/enum"

// ProgrammaticCreativeMedia 程序化创意素材信息
type ProgrammaticCreativeMedia struct {
	// ImageMode 创意素材类型
	ImageMode enum.MaterialMode `json:"image_mode,omitempty"`
	// VideoID 视频ID
	VideoID string `json:"video_id,omitempty"`
	// VideoCoverID 视频封面ID
	VideoCoverID string `json:"video_cover_id,omitempty"`
	// ImageIDs 图片ID列表
	ImageIDs []string `json:"image_ids,omitempty"`
	// CarouselID 素材库图文id
	CarouselID uint64 `json:"carousel_id,omitempty"`
	// AwemeCarouselID 抖音主页图文id
	AwemeCarouselID uint64 `json:"aweme_carousel_id,omitempty"`
	// IsAutoGenerate 是否为派生创意标识，1：是，0：不是
	IsAutoGenerate int `json:"is_auto_generate,omitempty"`
}
