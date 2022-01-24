package creative

import "github.com/bububa/oceanengine/marketing-api/enum/qianchuan"

// VideoMaterial 视频素材
type VideoMaterial struct {
	// ID 素材唯一标识
	ID uint64 `json:"id,omitempty"`
	// VideoID 视频ID
	VideoID string `json:"video_id,omitempty"`
	// VideoCoverID 视频封面ID
	VideoCoverID string `json:"video_cover_id,omitempty"`
	// AwemeItemID 抖音视频ID
	AwemeItemID uint64 `json:"aweme_item_id,omitempty"`
	// ImageMode 素材类型，见附录-枚举值
	ImageMode qianchuan.ImageMode `json:"image_mode,omitempty"`
	// IsAutoGenerate 是否为派生创意标识，1：是，0：不是
	IsAutoGenerate int `json:"is_auto_generate,omitempty"`
}
