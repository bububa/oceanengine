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
	// AwemeItemID 抖音视频ID
	AwemeItemID string `json:"aweme_item_id,omitempty"`
	// AwemeCarouselID 抖音主页图文id
	AwemeCarouselID uint64 `json:"aweme_carousel_id,omitempty"`
	// CarouselID 素材库图文id
	CarouselID uint64 `json:"carousel_id,omitempty"`
	// ImageIDs 图片ID列表
	ImageIDs []string `json:"image_ids,omitempty"`
	// IsAutoGenerate 是否为派生创意标识，1：是，0：不是
	IsAutoGenerate int `json:"is_auto_generate,omitempty"`
	// URL 视频地址，仅限同主体进行素材预览查看，若非同主体会返回“素材所属主体与开发者主体不一致无法获取URL”，链接1小时过期
	URL string `json:"url,omitempty"`
	// VideoPosterURL 视频首帧截图，仅限同主体进行素材预览查看，若非同主体会返回“素材所属主体与开发者主体不一致无法获取URL”，链接1小时过期
	VideoPosterURL string `json:"video_poster_url,omitempty"`
	// IamgeURL 图片地址
	ImageURL string `json:"iamge_url,omitempty"`
	// CarouselImages 图片信息
	CarouselImages []struct {
		// URL 图片预览url
		URL string `json:"url,omitempty"`
	} `json:"carousel_iamges,omitempty"`
	// CarouselAudio 音频信息
	CarouselAudio []struct {
		// URL 音频url
		URL string `json:"url,omitempty"`
	} `json:"carousel_audio,omitempty"`
	// CarouselDescription 图文描述
	CarouselDescription string `json:"carousel_description,omitempty"`
}
