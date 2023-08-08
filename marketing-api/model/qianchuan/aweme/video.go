package aweme

// Video 视频信息
type Video struct {
	// AwemeItemID 视频id
	AwemeItemID uint64 `json:"aweme_item_id,omitempty"`
	// AwemiItemTitle 视频标题
	AwemeItemTitle string `json:"aweme_item_title,omitempty"`
	// AwemeItemCover 视频封面
	AwemeItemCover string `json:"aweme_item_cover,omitempty"`
	// ItemType 视频类型
	// 1：图文视频
	// 0：普通视频
	ItemType int `json:"item_type,omitempty"`
	// VideoCoverURL 视频封面图片url
	VideoCoverURL string `json:"video_cover_url,omitempty"`
	// Width 视频宽度
	Width int `json:"width,omitempty"`
	// Height 视频高度
	Height int `json:"height,omitempty"`
	// URL 视频地址，链接1小时过期
	URL string `json:"url,omitempty"`
	// Duration 视频时长
	Duration float64 `json:"duration,omitempty"`
	// Title 抖音中的视频标题
	Title string `json:"title,omitempty"`
	// IsRecommend 是否推荐 :
	// 0: 不推荐
	// 1: 推荐
	IsRecommend int `json:"is_recommend,omitempty"`
	// ImgURL 图文预览链接
	ImgURL string `json:"img_url,omitempty"`
	// IsImg 是否图文视频
	// 1：是
	// 0：否
	IsImg int `json:"is_img,omitempty"`
	// 是否AI生成
	// false：不是AI生成
	// true：AI生成
	IsAiCreate bool `json:"is_ai_create,omitempty"`
}
