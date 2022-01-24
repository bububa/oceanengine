package creative

// AdData 广告计划数据
type AdData struct {
	// ThridIndustryID 三级行业ID
	ThirdIndustryID uint64 `json:"thrid_industry_id,omitempty"`
	// AdKeywords 创意标签。最多20个标签，且每个标签长度不超过10个字符
	AdKeywords []string `json:"ad_keywords,omitempty"`
	// Source 广告来源，4-20个字符，当推广目的为非应用下载或者应用下载且download_type为"EXTERNAL_URL时"时必填
	Source string `json:"source,omitempty"`
	// IesCoreUserID 品牌主页-推广抖音号，当传入此字段时表示开启抖音主页。广告视频将同步到您的主页下，在客户端点击广告头像将进入您的主页。创建后不可修改。
	IesCoreUserID string `json:"ies_core_user_id,omitempty"`
	// PlayableURL 搭配试玩素材URL，可通过【获取试玩素材列表】进行获取。
	PlayableURL string `json:"playable_url,omitempty"`
	// IsFeedAndFavSee 主页作品列表隐藏广告内容，默认值：0; 允选值：0（不隐藏），1（隐藏）
	IsFeedAndFavSee *int `json:"is_feed_and_fav_see,omitempty"`
	// CreativeAutoGenerateSwitch 是否开启自动派生创意，大通投时可填，默认值: 1允许值: 0(不启用), 1(启用)
	CreativeAutoGenerateSwitch *int `json:"creative_auto_generate_switch,omitempty"`
	// AppName 应用名，当广告计划的download_type为"DOWNLOAD_URL"时必填。1到40个字符，中文占2个字符
	AppName string `json:"app_name,omitempty"`
	// WebURL Android应用下载详情页（用户点击广告中“立即下载”按钮以外的区域时所到达的页面），当广告计划app_type为"APP_ANDROID"或快应用推广目的时, 必填; 可从此接口获取：【获取橙子建站站点列表】
	WebURL string `json:"web_url,omitempty"`
	// ExternalURL 落地页链接（支持橙子建站落地页）
	ExternalURL string `json:"external_url,omitempty"`
	// IsCommentDisabled 是否关闭评论，0为开启，1为关闭，默认值：0; 允许值: 0, 1
	IsCommentDisabled *int `json:"is_comment_disabled,omitempty"`
	// AdDownloadStatus 允许客户端下载视频功能，0为开启，即允许客户端下载视频；1为关闭，即不允许客户端下载视频。默认不传值，表示允许客户端下载视频。关闭客户端下载视频功能仅对本地上传的视频有效。
	AdDownloadStatus *int `json:"ad_download_status,omitempty"`
	// PriorityTrail 是否优先调起试玩。当推广目的为应用推广且使用搭配试玩素材时可以开启该功能。允许值：ON开启，OFF关闭
	PriorityTrail string `json:"priority_trail,omitempty"`
	// Supplements 云游戏
	Supplements []SupplementInfo `json:"supplements,omitempty"`
}

type SupplementInfo struct {
	// SupplementType 云游戏类型，允许值: CLOUD_GAME
	SupplementType string `json:"supplement_type,omitempty"`
	// Game 云游戏信息, 最多只允许填入一个
	Games []GameInfo `json:"game,omitempty"`
}

type GameInfo struct {
	// ID 云游戏id，对应【获取云游戏试玩素材列表】中的game_id字段
	ID string `json:"id,omitempty"`
	// Orientation 云游戏素材方向，允许值: VERTICAL竖屏，HORIZONTAL横屏
	Orientation string `json:"orientation,omitempty"`
}
