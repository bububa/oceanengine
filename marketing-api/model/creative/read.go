package creative

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// ReadRequest 创意详细信息API Request
type ReadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 计划ID
	AdID uint64 `json:"ad_id,omitempty"`
}

// Encode implement GetRequest interface
func (r ReadRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("ad_id", strconv.FormatUint(r.AdID, 10))
	return values.Encode()
}

// ReadResponse 创意详细信息API Response
type ReadResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CreativeDetail `json:"data,omitempty"`
}

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
	AdDownloadStatus int `json:"ad_download_status,omitempty"`
	// SmartInventory 是否使用优选广告位，0表示不使用优选，1表示使用，2表示标记该创意隶属的计划投放范围是通投智选
	SmartInventory int `json:"smart_inventory,omitempty"`
	// ComponentInfo 创意组件信息
	ComponentInfo []struct {
		// ComponentID 创意组件id
		ComponentID string `json:"component_id,omitempty"`
	} `json:"component_info,omitempty"`
	// SceneInventory 首选场景广告位，详见【附录-首选场景广告位】，使用首选场景广告位时默认忽略inventory_type字段，与scene_inventory不能同时传 允许值: "VIDEO_SCENE", "FEED_SCENE", "TAIL_SCENE"
	SceneInventory string `json:"scene_inventory,omitempty"`
	// CreativeMaterialMode 创意类型，该字段为STATIC_ASSEMBLE表示程序化创意，其他情况无该字段
	CreativeMaterialMode string `json:"creative_material_mode,omitempty"`
	// ProceduralPackageID 程序化创意包ID
	ProceduralPackageID uint64 `json:"procedural_package_id,omitempty"`
	// ProceduralPackageVersion 程序化创意包版本
	ProceduralPackageVersion uint64 `json:"procedural_package_version,omitempty"`
	// IsPresentedVideo 启用图片生成视频，允许值：0（不启用），1（启用）
	IsPresentedVideo int `json:"is_presented_video,omitempty"`
	// GenerateDerivedAd 是否开启衍生计划，1为开启，0为不开启
	GeneratedDerivedAd string `json:"generated_derived_ad,omitempty"`
	// ImageList 素材信息，程序化创意素材列表。最多包含12张图和6个视频。
	ImageList []struct {
		// ImageMode 素材类型
		ImageMode enum.ImageMode `json:"image_mode,omitempty"`
		// ImageID 视频图片ID
		ImageID string `json:"image_id,omitempty"`
		// VideoID 视频ID
		VideoID string `json:"video_id,omitempty"`
		// ImageIDs 图片ID列表
		ImageIDs []string `json:"image_ids,omitempty"`
		// TemplateIDs 模版ID列表
		TemplateIDs []uint64 `json:"template_ids,omitempty"`
	} `json:"image_list,omitempty"`
	// TitleList 标题信息，程序化创意标题列表。最多包含10个标题
	TitleList []struct {
		// Title 创意标题
		Title string `json:"title,omitempty"`
		// CreativeWordIDs 动态词包ID，可使用动态词包查询接口查询数据
		CreativeWordIDs []uint64 `json:"creative_word_ids,omitempty"`
		// DpaWordIDs DPA词包ID列表，针对DPA广告
		DpaWordIDs []uint64 `json:"dpa_word_ids,omitempty"`
	} `json:"title_list,omitempty"`
	// Creatives 素材信息, 首选投放位置和创意类型决定素材规格。程序化创意只有在审核通过后才有值
	Creatives []Creative `json:"creative,omitempty"`
	// Source 广告来源
	Source string `json:"source,omitempty"`
	// IesCoreUserID 广告主绑定的抖音ID
	IesCoreUserID string `json:"ies_core_user_id,omitempty"`
	// IsFeedAndFavSee 是否隐藏抖音主页，0：不隐藏，1：隐藏
	IsFeedAndFavSee int `json:"is_feed_and_fav_see,omitempty"`
	// CreativeAutoGenerateSwitch 是否开启自动生成素材，delivery_range为UNIVERSAL：通投智选时返回，0：不启用,1：启用
	CreativeAutoGenerateSwitch int `json:"creative_auto_generate_switch,omitempty"`
	// AppName 应用名
	AppName string `json:"app_name,omitempty"`
	// SubTitle APP 副标题。
	SubTitle string `json:"sub_title,omitempty"`
	// WebUrl Android应用下载详情页
	WebUrl string `json:"web_url,omitempty"`
	// ActionText 行动号召
	ActionText string `json:"action_text,omitempty"`
	// PlayableUrl 试玩素材URL
	PlayableUrl string `json:"playable_url,omitempty"`
	// IsCommentDisable 是否关闭评论
	IsCommentDisable int `json:"is_comment_disable,omitempty"`
	// CloseVideoDetail 是否关闭视频详情页落地页(勾选该选项后,视频详情页中不默认弹出落地页,仅对视频广告生效)
	CloseVideoDetail int `json:"close_video_detail,omitempty"`
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
	VideoPlayEffectTrackUrl []string `json:"video_play_effect_track_url,omitempty"`
	// VideoPlayDoneTrackUrl 视频播完（监测链接）
	VideoPlayDoneTrackUrl []string `json:"video_play_done_track_url,omitempty"`
	// VideoPlayTrackUrl 视频播放（监测链接）
	VideoPlayTrackUrl []string `json:"video_play_track_url,omitempty"`
	// TrackUrlSendType 数据发送方式; 允许值: SERVER_SEND(服务器端上传), CLIENT_SEND(客户端上传)
	TrackUrlSendType string `json:"track_url_send_type,omitempty"`
	// PromotionCard 商品推广卡片
	PromotionCard *PromotionCard `json:"promotion_card,omitempty"`
	// CollocationType 云游戏类型，允许值"CLOUD_GAME"云游戏
	CollocationType string `json:"collocation_type,omitempty"`
	// Supplements 云游戏素材
	Supplements []Supplement `json:"supplements,omitempty"`
}

// Supplement 云游戏素材
type Supplement struct {
	// GameID 云游戏id
	GameID string `json:"game_id,omitempty"`
	// Orientation 云游戏横竖屏，"VERTICAL"竖屏, "HORIZONTAL"横屏
	Orientation string `json:"orientation,omitempty"`
}
