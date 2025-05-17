package creative

import "github.com/bububa/oceanengine/marketing-api/enum"

// AdData 广告计划数据
type AdData struct {
	// ThridIndustryID 三级行业ID
	ThirdIndustryID uint64 `json:"thrid_industry_id,omitempty"`
	// ParamsType 链接类型
	ParamsType string `json:"params_type,omitempty"`
	// DpaExternalUrlField 落地页链接字段选择
	DpaExternalUrlField *string `json:"dpa_external_url_field,omitempty"`
	// AdKeywords 创意标签。最多20个标签，且每个标签长度不超过10个字符
	AdKeywords []string `json:"ad_keywords,omitempty"`
	// Source 广告来源，4-20个字符，当推广目的为非应用下载或者应用下载且download_type为"EXTERNAL_URL时"时必填
	Source string `json:"source,omitempty"`
	// EnableSmartSource 是否开启来源智能生成，允许值：ON 开启，OFF 关闭
	EnableSmartSource enum.OnOff `json:"enable_smart_source,omitempty"`
	// IesCoreUserID 品牌主页-推广抖音号，当传入此字段时表示开启抖音主页。广告视频将同步到您的主页下，在客户端点击广告头像将进入您的主页。创建后不可修改。
	IesCoreUserID string `json:"ies_core_user_id,omitempty"`
	// IsPresentedVideo 自动生成视频素材，利用已上传的图片与视频生成更多优质的短视频素材：1（启用），0（不启用）默认值: 0
	IsPresentedVideo *int `json:"is_presented_video,omitempty"`
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
	// ExternalUrlParams 落地页检测参数
	ExternalUrlParams *string `json:"external_url_params,omitempty"`
	// IsCommentDisable 是否关闭评论，0为开启，1为关闭，默认值：0; 允许值: 0, 1
	IsCommentDisable *int `json:"is_comment_disable,omitempty"`
	// AdDownloadStatus 允许客户端下载视频功能，0为开启，即允许客户端下载视频；1为关闭，即不允许客户端下载视频。默认不传值，表示允许客户端下载视频。关闭客户端下载视频功能仅对本地上传的视频有效。
	AdDownloadStatus *int `json:"ad_download_status,omitempty"`
	// PriorityTrail 是否优先调起试玩。当推广目的为应用推广且使用搭配试玩素材时可以开启该功能。允许值：ON开启，OFF关闭
	PriorityTrail enum.OnOff `json:"priority_trail,omitempty"`
	// Supplements 云游戏
	Supplements []SupplementInfo `json:"supplements,omitempty"`
	// DynamicCreativeSwitch 启用动态创意类型,详见【附录-动态创意类型】
	// 允许值:DYNAMIC_CREATIVE_TITLE, DYNAMIC_CREATIVE_ABSTRACT，DYNAMIC_CREATIVE_SUBLINK，DYNAMIC_CREATIVE_ON，默认DYNAMIC_CREATIVE_ON当传入不为空时，等同于传入DYNAMIC_CREATIVE_ON启用动态创意，当传入[]时，关闭动态创意
	// 不传时，不改变已有的值
	// 注意：该字段为【增量更新】
	DynamicCreativeSwitch []enum.DynamicCreativeType `json:"dynamic_creative_switch,omitempty"`
	// OpenURL 直达链接，只在电商店铺推广推广目的下有效
	OpenURL string `json:"open_url,omitempty"`
	// MiniProgramInfo 字节小程序信息
	MiniProgramInfo *MiniProgramInfo `json:"mini_program_info,omitempty"`
	// AnchorRelatedType 原生锚点启用类型，允许值：不启用OFF，自动生成AUTO，手动选择SELECT
	// 默认值为OFF
	// 自动生成AUTO仅应用推广目的下时支持
	AnchorRelatedType string `json:"anchor_related_type,omitempty"`
	// AnchorType 锚点类型，允许值：
	// - 应用下载-游戏：APP_GAME
	// - 应用下载-网服：APP_INTERNET_SERVICE
	// - 应用下载-电商：APP_SHOP
	// - 高级在线预约：ONLINE_SUBSCRIBE
	// 当 anchor_related_type = SELECT时必填
	AnchorType enum.AnchorType `json:"anchor_type,omitempty"`
	// AnchorID 原生锚点id，当 anchor_related_type = SELECT时必填，可从【获取账户下原生锚点】接口中获取
	AnchorID string `json:"anchor_id,omitempty"`
}

// SupplementInfo 云游戏列表
type SupplementInfo struct {
	// SupplementType 云游戏类型，允许值: CLOUD_GAME
	SupplementType string `json:"supplement_type,omitempty"`
	// Game 云游戏信息, 最多只允许填入一个
	Games []GameInfo `json:"games,omitempty"`
}

// GameInfo 云游戏信息
type GameInfo struct {
	// ID 云游戏id，对应【获取云游戏试玩素材列表】中的game_id字段
	ID string `json:"id,omitempty"`
	// Orientation 云游戏素材方向，允许值: VERTICAL竖屏，HORIZONTAL横屏
	Orientation string `json:"orientation,omitempty"`
}

// MiniProgramInfo 字节小程序信息
type MiniProgramInfo struct {
	// AppID 小程序/小游戏id
	AppID string `json:"app_id,omitempty"`
	// StartPath 启动路径，小程序类型必传，小游戏类型不传值
	StartPath string `json:"start_path,omitempty"`
	// Params 页面监测参数
	Params string `json:"params,omitempty"`
	// Type 小程序类型，当使用 mini_program_info 时，该字段必填
	// 允许值：BYTE_GAME 小游戏、BYTE_APP 小程序
	Type enum.MiniProgramType `json:"type,omitempty"`
	// URL 字节小程序调起链接
	URL string `json:"url,omitempty"`
}
