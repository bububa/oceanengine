package creative

import "github.com/bububa/oceanengine/marketing-api/enum"

// CreativeDetail 创意详情
type CreativeDetail struct {
	// AdID 广告计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ModifyTime 时间戳，用于判断创意版本
	ModifyTime string `json:"modify_time,omitempty"`
	// InventoryType 创意首选投放位置,详见【附录-首选投放位置】，请注意广告位中的信息流（INVENTORY_FEED）与详情页（INVENTORY_TEXT_LINK）已合并为头条系（依然使用INVENTORY_FEED字段，名称更改为头条系）
	InventoryType []enum.StatInventoryType `json:"inventory_type,omitempty"`
	// ExternalUrl 落地页链接，新版营销链路下创意支持
	ExternalUrl string `json:"external_url,omitempty"`
	// AdDownloadStatus 允许客户端下载视频功能，0为开启，即允许客户端下载视频；1为关闭，即不允许客户端下载视频，该字段为空与0效果一致，即表示允许客户端下载视频。关闭客户端下载视频功能仅对本地上传的视频有效
	AdDownloadStatus *int `json:"ad_download_status,omitempty"`
	// SmartInventory 是否使用优选广告位，0表示不使用优选，1表示使用，2表示标记该创意隶属的计划投放范围是通投智选
	SmartInventory int `json:"smart_inventory,omitempty"`
	// ComponentInfo 创意组件信息
	ComponentInfo []ComponentMaterial `json:"component_info,omitempty"`
	// SceneInventory 首选场景广告位，详见【附录-首选场景广告位】，使用首选场景广告位时默认忽略inventory_type字段，与scene_inventory不能同时传 允许值: "VIDEO_SCENE", "FEED_SCENE", "TAIL_SCENE"
	SceneInventory string `json:"scene_inventory,omitempty"`
	// CreativeMaterialMode 创意类型，该字段为STATIC_ASSEMBLE表示程序化创意，其他情况无该字段
	CreativeMaterialMode enum.CreativeMaterialMode `json:"creative_material_mode,omitempty"`
	// ProceduralPackageID 程序化创意包ID
	ProceduralPackageID uint64 `json:"procedural_package_id,omitempty"`
	// ProceduralPackageVersion 程序化创意包版本
	ProceduralPackageVersion uint64 `json:"procedural_package_version,omitempty"`
	// IsPresentedVideo 启用图片生成视频，允许值：0（不启用），1（启用）
	IsPresentedVideo *int `json:"is_presented_video,omitempty"`
	// GenerateDerivedAd 是否开启衍生计划，1为开启，0为不开启
	GeneratedDerivedAd string `json:"generated_derived_ad,omitempty"`
	// ImageList 素材信息，程序化创意素材列表。最多包含12张图和6个视频。
	ImageList []ImageInfo `json:"image_list,omitempty"`
	// TitleList 标题信息，程序化创意标题列表。最多包含10个标题
	TitleList []TitleMaterial `json:"title_list,omitempty"`
	// AbstractList 搜索广告字段
	AbstractList []AbstractMaterial `json:"abstract_list,omitempty"`
	// Creatives 素材信息, 首选投放位置和创意类型决定素材规格。程序化创意只有在审核通过后才有值
	Creatives []Creative `json:"creatives,omitempty"`
	// Source 广告来源
	Source string `json:"source,omitempty"`
	// IesCoreUserID 广告主绑定的抖音ID
	IesCoreUserID string `json:"ies_core_user_id,omitempty"`
	// IsFeedAndFavSee 是否隐藏抖音主页，0：不隐藏，1：隐藏
	IsFeedAndFavSee *int `json:"is_feed_and_fav_see,omitempty"`
	// CreativeAutoGenerateSwitch 是否开启自动生成素材，delivery_range为UNIVERSAL：通投智选时返回，0：不启用,1：启用
	CreativeAutoGenerateSwitch *int `json:"creative_auto_generate_switch,omitempty"`
	// AppName 应用名
	AppName string `json:"app_name,omitempty"`
	// SubTitle APP 副标题。
	SubTitle string `json:"sub_title,omitempty"`
	// WebUrl Android应用下载详情页
	WebUrl string `json:"web_url,omitempty"`
	// OpenUrl 直达链接，只在电商店铺推广推广目的下有效
	OpenUrl string `json:"open_url,omitempty"`
	// ActionText 行动号召
	ActionText string `json:"action_text,omitempty"`
	// PlayableUrl 试玩素材URL
	PlayableUrl string `json:"playable_url,omitempty"`
	// IsCommentDisable 是否关闭评论
	IsCommentDisable *int `json:"is_comment_disable,omitempty"`
	// CloseVideoDetail 是否关闭视频详情页落地页(勾选该选项后,视频详情页中不默认弹出落地页,仅对视频广告生效)
	CloseVideoDetail *int `json:"close_video_detail,omitempty"`
	// CreativeDisplayMode 创意展现方式
	CreativeDisplayMode enum.CreativeDisplayMode `json:"creative_display_mode,omitempty"`
	// AdvancedCreativeType 附加创意类型
	AdvancedCreativeType enum.AdvancedCreativeType `json:"advanced_creative_type,omitempty"`
	// AdvancedCreativeTitle 附加创意副标题
	AdvancedCreativeTitle string `json:"advanced_creative_title,omitempty"`
	// PhoneNumber 电话号码(当附加创意类型为ATTACHED_CREATIVE_PHONE时返回)
	PhoneNumber string `json:"phone_number,omitempty"`
	// ButtonText 按钮文本(当附加创意类型不为ATTACHED_CREATIVE_NONE时返回)
	ButtonText string `json:"button_text,omitempty"`
	// FormUrl 表单提交链接(当附加创意类型为ATTACHED_CREATIVE_FORM时返回)
	FormUrl string `json:"form_url,omitempty"`
	// CommmerceCards 产品（商业卡）信息。如果没有启用，那么不返回相关字段。
	CommerceCards []CommerceCard `json:"commerce_cards,omitempty"`
	// ThirdIndustryID 三级行业ID
	ThirdIndustryID uint64 `json:"third_industry_id,omitempty"`
	// AdKeywords 创意标签
	AdKeywords []string `json:"ad_keywords,omitempty"`
	// TrackUrl 展示（监测链接）
	TrackUrl string `json:"track_url,omitempty"`
	// ActionTrackUrl 点击（监测链接）（当推广目的为应用下载且创建计划传递了convert_id，系统会自动获取转化中的点击监测链接，且不可修改）
	ActionTrackUrl string `json:"action_track_url,omitempty"`
	// VideoPlayEffectTrackUrl 视频有效播放（监测链接）
	VideoPlayEffectTrackUrl string `json:"video_play_effect_track_url,omitempty"`
	// VideoPlayDoneTrackUrl 视频播完（监测链接）
	VideoPlayDoneTrackUrl string `json:"video_play_done_track_url,omitempty"`
	// VideoPlayTrackUrl 视频播放（监测链接）
	VideoPlayTrackUrl string `json:"video_play_track_url,omitempty"`
	// TrackUrlSendType 数据发送方式; 允许值: SERVER_SEND(服务器端上传), CLIENT_SEND(客户端上传)
	TrackUrlSendType string `json:"track_url_send_type,omitempty"`
	// PromotionCard 商品推广卡片
	PromotionCard *PromotionCard `json:"promotion_card,omitempty"`
	// CollocationType 云游戏类型，允许值"CLOUD_GAME"云游戏
	CollocationType string `json:"collocation_type,omitempty"`
	// Supplements 云游戏素材
	Supplements []Supplement `json:"supplements,omitempty"`
	// IsSmartTitle
	IsSmartTitle *int `json:"is_smart_title,omitempty"`
	// AdCategory
	AdCategory uint64 `json:"ad_category,omitempty"`
	// PriorityTrial 是否优先调起试玩。当推广目的为应用推广且使用搭配试玩素材是可以开启该功能。允许值：ON开启，OFF关闭，默认关闭
	PriorityTrial enum.OnOff `json:"priority_trial,omitempty"`
	// DynamicCreativeSwitch 启用动态创意类型,详见【附录-动态创意类型】
	// 允许值:DYNAMIC_CREATIVE_TITLE, DYNAMIC_CREATIVE_ABSTRACT，DYNAMIC_CREATIVE_SUBLINK，DYNAMIC_CREATIVE_ON，默认DYNAMIC_CREATIVE_ON当传入不为空时，等同于传入DYNAMIC_CREATIVE_ON启用动态创意，当传入[]时，关闭动态创意
	// 不传时，不改变已有的值
	// 注意：该字段为【增量更新】
	DynamicCreativeSwitch []enum.DynamicCreativeType `json:"dynamic_creative_switch,omitempty"`
	// MiniProgramInfo 字节小程序信息
	MiniProgramInfo *MiniProgramInfo `json:"mini_program_info,omitempty"`
}

// Supplement 云游戏素材
type Supplement struct {
	// GameID 云游戏id
	GameID string `json:"game_id,omitempty"`
	// Orientation 云游戏横竖屏，"VERTICAL"竖屏, "HORIZONTAL"横屏
	Orientation string `json:"orientation,omitempty"`
}
