package creative

import "github.com/bububa/oceanengine/marketing-api/enum"

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
	ImageMode enum.MaterialMode `json:"image_mode,omitempty"`
	// IsAutoGenerate 是否为派生创意标识，1：是，0：不是
	IsAutoGenerate int `json:"is_auto_generate,omitempty"`
	// URL 视频地址，仅限同主体进行素材预览查看，若非同主体会返回“素材所属主体与开发者主体不一致无法获取URL”，链接1小时过期
	URL string `json:"url,omitempty"`
	// PosterURL 视频首帧截图，仅限同主体进行素材预览查看，若非同主体会返回“素材所属主体与开发者主体不一致无法获取URL”，链接1小时过期
	PosterURL string `json:"poster_url,omitempty"`
}
