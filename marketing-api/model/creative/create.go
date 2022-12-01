package creative

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建广告创意 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 广告计划ID，计划ID要属于广告主ID，且非删除计划，否则会报错
	AdID uint64 `json:"ad_id,omitempty"`
	// InventoryType 广告投放位置（首选媒体）
	InventoryType enum.StatInventoryType `json:"inventory_type,omitempty"`
	// SmartInventory 优选广告位，0表示不使用优选，1表示使用，使用优选广告位的时候默认忽略inventory_type字段;默认值: 0; 允许值: 0、1
	SmartInventory int `json:"smart_inventory,omitempty"`
	// CreativeMaterialMode 创意方式，当值为"STATIC_ASSEMBLE"表示程序化创意，其他情况不传字段
	CreativeMaterialMode enum.CreativeMaterialMode `json:"creative_material_mode,omitempty"`
	// ProceduralPackageID 程序化创意包ID，可通过【查询程序化创意包】接口进行查询，程序化创意包与自定义素材（title_list和image_list）不能同时使用，否则会报错;仅支持程序化创意，头条文章、DPA推广类型暂不支持
	ProceduralPackageID uint64 `json:"procedural_package_id,omitempty"`
	// IsPresentedVideo 自动生成视频素材，利用已上传的图片与视频生成更多优质的短视频素材：1（启用），0（不启用）默认值: 0
	IsPresentedVideo int `json:"is_presented_video,omitempty"`
	// ImageList 素材信息，creative_material_mode为"STATIC_ASSEMBLE"时必填，字段说明见下表。最多包含12张图和10个视频。
	ImageList []ImageInfo `json:"image_list,omitempty"`
	// TitleList 标题信息，creative_material_mode为"STATIC_ASSEMBLE"时必填，字段说明见下表。最多包含10个标题。
	TitleList []TitleMaterial `json:"title_list,omitempty"`
	// ComponentInfo 组件信息
	ComponentInfo []ComponentMaterial `json:"component_info,omitempty"`
	// Creatives 自定义素材信息, 最多支持10个创意。首选投放位置和创意类型决定素材规格。当为程序化创意时，该字段不填数据，值为[]
	Creatives []Creative `json:"creatives,omitempty"`
	// Source 广告来源，4-20个字符，当推广目的为非应用下载或者应用下载且download_type为"EXTERNAL_URL时"时必填
	Source string `json:"source,omitempty"`
	// IesCoreUserID 品牌主页-推广抖音号，当传入此字段时表示开启抖音主页。广告视频将同步到您的主页下，在客户端点击广告头像将进入您的主页。创建后不可修改。
	IesCoreUserID string `json:"ies_core_user_id,omitempty"`
	// IsFeedAndFavSee 主页作品列表隐藏广告内容，默认值：0; 允选值：0（不隐藏），1（隐藏）
	IsFeedAndFavSee *int `json:"is_feed_and_fav_see,omitempty"`
	// CreativeAutoGenerateSwitch 是否开启自动派生创意，大通投时可填，默认值: 1允许值: 0(不启用), 1(启用)
	CreativeAutoGenerateSwitch *int `json:"creative_auto_generate_switch,omitempty"`
	// AppName 应用名，当广告计划的download_type为"DOWNLOAD_URL"时必填。1到40个字符，中文占2个字符
	AppName string `json:"app_name,omitempty"`
	// SubTitle APP 副标题。仅推广目标为APP，4到24个字符，填写Android下载链接时可设置
	SubTitle string `json:"sub_title,omitempty"`
	// WebURL Android应用下载详情页（用户点击广告中“立即下载”按钮以外的区域时所到达的页面），当广告计划app_type为"APP_ANDROID"或快应用推广目的时, 必填; 可从此接口获取：【获取橙子建站站点列表】
	WebURL string `json:"web_url,omitempty"`
	// ActionText 行动号召（仅应用下载推广类型有效）;备注：应用下载的行动号召字段使用action_text，门店与销售线索行动号召使用button_text;请求值可从接口【行动号召字段内容获取】进行获取，如果不传参默认为立即下载
	ActionText string `json:"action_text,omitempty"`
	// PlayableURL 搭配试玩素材URL，可通过【获取试玩素材列表】进行获取。
	PlayableURL string `json:"playable_url,omitempty"`
	// IsCommentDisabled 是否关闭评论，0为开启，1为关闭，默认值：0; 允许值: 0, 1
	IsCommentDisabled *int `json:"is_comment_disabled,omitempty"`
	// PromotionCard 商品推广卡片，如不传，则创意中没有推广卡片
	PromotionCard []PromotionCard `json:"promotion_card,omitempty"`
	// AdDownloadStatus 允许客户端下载视频功能，0为开启，即允许客户端下载视频；1为关闭，即不允许客户端下载视频。默认不传值，表示允许客户端下载视频。关闭客户端下载视频功能仅对本地上传的视频有效。
	AdDownloadStatus *int `json:"ad_download_status,omitempty"`
	// AdvancedCreativeType 附加创意类型。直播创意枚举：ATTACHED_CREATIVE_LIVE_CARD（直播卡片）
	AdvancedCreativeType enum.AdvancedCreativeType `json:"advanced_creative_type,omitempty"`
	// AdvancedCreativeTitle 副标题，最多24个字符
	AdvancedCreativeTitle string `json:"advanced_creative_title,omitempty"`
	// PhoneNumber 电话号码。当附加创意类型为"ATTACHED_CREATIVE_PHONE"时必填
	PhoneNumber string `json:"phone_number,omitempty"`
	// ButtonText 按钮文本，即行动号召，当附加创意类型非"ATTACHED_CREATIVE_NONE"时填写，请求值可从接口【行动号召字段内容获取】进行获取
	ButtonText string `json:"button_text,omitempty"`
	// FormURL 表单提交链接。当附加创意类型为"ATTACHED_CREATIVE_FORM"时 必填，必须为今日头条建站地址：【查询已有表单列表】
	FormURL string `json:"form_url,omitempty"`
	// CommerceCards 产品（商业卡）信息;目前为白名单功能，如需使用请联系平台运营
	CommerceCards []CommerceCard `json:"commerce_cards,omitempty"`
	// TrackURL 展示（监测链接）
	TrackURL string `json:"track_url,omitempty"`
	// ActionTrackURL 点击（监测链接）（当推广目的为应用下载且创建计划传递了convert_id，系统会自动获取转化中的点击监测链接，且不可修改）
	ActionTrackURL string `json:"action_track_url,omitempty"`
	// VideoPlayEffectiveTrackURL 视频有效播放（监测链接），投放范围为穿山甲时暂不支持设置此链接
	VideoPlayEffectiveTrackURL string `json:"video_play_effective_track_url,omitempty"`
	// VideoPlayDoneTrackURL 视频播完（监测链接），投放范围为穿山甲时暂不支持设置此链接
	VideoPlayDoneTrackURL string `json:"video_play_done_track_url,omitempty"`
	// VideoPlayTrackURL 视频播放（监测链接），投放范围为穿山甲时暂不支持设置此链接
	VideoPlayTrackURL string `json:"video_play_track_url,omitempty"`
	// TrackURLSendType 数据发送方式，不可修改,默认值: SERVER_SEND; 允许值: SERVER_SEND(服务器端上传), CLIENT_SEND(客户端上传);客户端上传是指由客户端直接上报给监测平台的服务器, 只有白名单用户才可使用CLIENT_SEND(客户端上传), 如果需要开通请找对接的销售、运营
	TrackURLSendType string `json:"track_url_send_type,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建广告创意 API Response
type CreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CreateResponseData `json:"data,omitempty"`
}

// CreateResponseData json返回值
type CreateResponseData struct {
	CreateRequest
	// ProceduralPackageVersion 程序化创意包版本
	ProceduralPackageVersion string `json:"procedural_package_version,omitempty"`
	// GenerateDerivedAd 是否开启衍生计划，1为开启，0为不开启; 默认值: 0
	GenerateDerivedAd int `json:"generate_derived_ad,omitempty"`
	// CloseVideoDetail 是否关闭视频详情页落地页(勾选该选项后,视频详情页中不默认弹出落地页,仅对视频广告生效); 允许值: 0, 1
	CloseVideoDetail int `json:"close_video_detail,omitempty"`
	// CreativeDisplayMode 创意展现方式
	CreativeDisplayMode enum.CreativeDisplayMode `json:"creative_display_mode,omitempty"`
}
