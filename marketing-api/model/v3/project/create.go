package project

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建项目 API Request
type CreateRequest struct {
	// AdvertiserID 广告账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Operation 计划状态，允许值: ENABLE开启（默认值）,DISABLE关闭
	Operation enum.OptStatus `json:"operation,omitempty"`
	// DeliveryMode 投放模式，允许值：
	// MANUAL手动投放(默认值）、PROCEDURAL自动投放
	// 自动投放仅支持landing_type=APP或MICRO_GAME或LINK
	// 自动投放不支持SEARCH搜索广告
	// 当marketing_goal= LIVE时，仅支持MANUAL手动投放
	DeliveryMode enum.DeliveryMode `json:"delivery_mode,omitempty"`
	// LandingType 推广目的，允许值：APP 应用推广、LINK 销售线索推广、MICRO_GAME 小程序、SHOP 电商店铺推广、QUICK_APP快应用、NATIVE_ACTION 原生互动、DPA商品目录
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// AppPromotionType 子目标，当 landing_type = APP 有效且必填
	// 允许值：DOWNLOAD 应用下载、LAUNCH 应用调用、RESERVE 预约下载
	// 当delivery_mode = PROCEDURAL 时仅支持DOWNLOAD应用下载；
	// 当marketing_goal= LIVE时，仅支持DOWNLOAD应用下载、LAUNCH 应用调起
	AppPromotionType enum.AppPromotionType `json:"app_promotion_type,omitempty"`
	// MarketingGoal 营销场景，允许值：VIDEO_AND_IMAGE 短视频/图片，LIVE直播,
	// LIVE仅支持已在广告平台签署直播推广协议的账户，支持的landing_type有应用/小程序/线索/原生互动
	// 当delivery_mode选择PROCEDURAL且landing_type选择LINK时，仅支持VIDEO_AND_IMAGE
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// 广告类型，允许值：ALL 通投广告 SEARCH 搜索广告
	// 当 marketing_goal= LIVE时，仅支持ALL
	// 仅当landing_type=APP/LINk&&delivery_mode=MANUAL时支持搜索广告，否则报错
	AdType enum.CampaignType `json:"ad_type,omitempty"`
	// DeliveryType 投放类型，不传默认为NORMAL常规投放，允许值：
	// NORMAL 常规投放（默认值）
	// DURATION 周期稳投（目前仅支持搜索广告）
	// 当前仅支持搜索广告，即ad_type = SEARCH下传入该枚举值有效，否则报错
	// 当landing_type = APP 应用推广、LINK 销售线索推广、MICRO_GAME 小程序时，允许创建周期稳投搜索广告
	DeliveryType enum.DeliveryType `json:"delivery_type,omitempty"`
	// Name 项目名称
	Name string `json:"name,omitempty"`
	// SearchBidRatio 出价系数，默认系数为1，出价系数可通过【获取快投推荐出价系数】查询，小数点后最多两位,取值范围 [1,2]
	// 当符合以下所有条件时填写有效
	// 1. bid_type != NO_BID && pricing = PRICING_OCPM
	// 2. deep_bid_type = DEEP_BID_DEFAULT 无深度优化方式 /BID_PER_ACTION 每次付费
	SearchBidRatio float64 `json:"search_bid_ratio,omitempty"`
	// AudienceExtend 定向拓展, 允许值：ON:开启（默认值）， OFF:关闭
	AudienceExtend enum.OnOff `json:"audience_extend,omitempty"`
	// Keywords 搜索关键词列表
	Keywords []Keyword `json:"keywords,omitempty"`
	// AutoTraficExtend 智能拓流 ，允许值：ON开启； OFF关闭
	// 仅支持ad_type = SEARCH时设置
	// 自定义关键词和智能拓流二者必须开启一个：若keywords为空，智能拓流 auto_extend_traffic需为ON
	// 对于搜索极速智投项目，若设置blue_flow_keyword_name蓝海关键词，智能拓流默认值为ON，且不得设置为OFF
	AutoExtendTraffic enum.OnOff `json:"auto_extend_traffic,omitempty"`
	// BlueFlowPackage 搜索蓝海流量投放相关参数
	BlueFlowPackage *BlueFlowPackage `json:"blue_flow_package,omitempty"`
	// RelatedProduct 关联产品投放相关
	RelatedProduct *RelatedProduct `json:"related_product,omitempty"`
	// DpaCategories 商品投放范围，分类列表，由【DPA商品广告-获取DPA分类】 得到
	// 个数限制 [0, 1000]
	// 不传和传空数组即为不限商品投放范围
	// DPA推广目的下有效
	DpaCategories []uint64 `json:"dpa_categories,omitempty"`
	// DpaProductTarget 自定义筛选条件（商品投放条件）。用于圈定商品投放范围，结合商品库字段搭配判断条件，圈定商品投放范围。获取商品库元信息-商品广告-商业开放平台
	// 数组长度限制：最大5条
	// DPA推广目的下有效
	DpaProductTarget []DpaProductTarget `json:"dpa_product_target,omitempty"`
	// DownloadURL 下载链接
	DownloadURL string `json:"download_url,omitempty"`
	// AppName 应用名称，dpa_adtype = DPA_APP时、dpa_adtype = DPA_APP 时必填
	AppName string `json:"app_name,omitempty"`
	// DownloadType 下载方式，枚举值：DOWNLOAD_URL 直接下载、EXTERNAL_URL 落地页下载
	DownloadType enum.DownloadType `json:"download_type,omitempty"`
	// DownloadMode 优先从系统应用商店下载（下载模式），枚举值：APP_STORE_DELIVERY 优先商店下载、 DEFAULT 默认下载
	DownloadMode enum.DownloadMode `json:"download_mode,omitempty"`
	// QuickAppId 快应用资产id ，从【查询快应用信息】接口获取，仅支持已通过审核的快应用资产
	QuickAppId uint64 `json:"quick_app_id,omitempty"`
	// LaunchType 调起方式，枚举值： DIRECT_OPEN 直接调起、EXTERNAL_OPEN 落地页调起
	LaunchType enum.LaunchType `json:"launch_type,omitempty"`
	// OpenURL Deeplink直达链接，landing_type = APP 且子目标为 LAUNCH 时有效且必填
	// 直达链接仅支持部分App唤起（点击唤起APP），点击创意将优先跳转App，再根据投放内容跳转相关链接
	OpenURL string `json:"open_url,omitempty"`
	// UlinkURL ulink直达链接，landing_type = APP 且子目标为LAUNCH 时有效 仅支持穿山甲广告位
	UlinkURL string `json:"ulink_url,omitempty"`
	// SubscribeURL 预约下载链接，landing_type = APP 且子目标为 RESERVE 时有效且必填
	SubscribeURL string `json:"subscribe_url,omitempty"`
	// AssetType 资产类型 landing_type = LINK 或SHOP时有效且必填
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// MultiAssetType 多投放载体
	// 允许值：ORANGE_AND_AWEME优选投放橙子落地页和抖音主页
	// 同时满足以下条件支持设置多投放载体：landing_type = LINK && marketing_goal = VIDEO_AND_IMAGE
	// 注意：同时传入asset_type 和 multi_asset_type 实际会生效多投放载体投放
	MultiAssetType enum.MultiDeliveryMedium `json:"multi_asset_type,omitempty"`
	// 小程序类型，landing_type = MICRO_GAME 时有效且必填
	// 允许值： WECHAT_GAME 微信小游戏、WECHAT_APP微信小程序、BYTE_GAME字节小游戏、BYTE_APP字节小程序
	MicroPromotionType enum.MicroPromotionType `json:"micro_promotion_type,omitempty"`
	// DpaAdType DPA广告类型，
	// 允许值: DPA_LINK 落地页
	// 当landing_type为dpa时有效且必填
	DpaAdType enum.DpaAdType `json:"dpa_adtype,omitempty"`
	// 字节小程序/小游戏资产id，通过 工具-字节小程序 接口获取
	// 10.30日后上线条件必填逻辑
	// landing_type = MICRO_GAME下，且micro_promotion_type=BYTE_GAME字节小游戏、BYTE_APP字节小程序，必填
	// landing_type=
	// LINK SHOP下，若选择事件资产为字节小程序类型，则必填
	MicroAppInstanceID uint64 `json:"micro_app_instance_id,omitempty"`
	// OptimizeGoal 优化目标
	OptimizeGoal *OptimizeGoal `json:"optimize_goal,omitempty"`
	// LandingPageStayTime 店铺停留时长，单位为毫秒，当external_action为AD_CONVERT_TYPE_STAY_TIME时有效且必填
	LandingPageStayTime enum.LandingPageStayTime `json:"landing_page_stay_time,omitempty"`
	// DeliveryRange 广告版位
	DeliveryRange *DeliveryRange `json:"delivery_range,omitempty"`
	// Audience 定向设置
	Audience *Audience `json:"audience,omitempty"`
	// DeliverySetting 投放设置
	DeliverySetting *DeliverySetting `json:"delivery_setting,omitempty"`
	// TrackURLSetting 监测链接设置
	TrackURLSetting *TrackURLSetting `json:"track_url_setting,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建项目 API Response
type CreateResponse struct {
	model.BaseResponse
	Data *CreateResult `json:"data,omitempty"`
}

type CreateResult struct {
	// ProjectID 项目id
	ProjectID uint64 `json:"project_id,omitempty"`
	// SuppleMentaryAgreementInfo 星广联投投放协议查看地址，仅短剧行业投放星广联投项目时会返回。建议您在开启项目投放前协同广告主仔细确认协议内容后再进行投放
	// https://open.oceanengine.com/labels/7/docs/1815754527706187
	SuppleMentaryAgreementInfo string `json:"supple_mentary_agreement_info,omitempty"`
	// ErrorKeywordList 上传失败的关键词list
	ErrorKeywordList []ErrorKeyword `json:"error_keyword_list,omitempty"`
	// AutoDurationProjectCount 账户下剩余可建智能托管项目额度数量（UBA智能托管项目创建成功时返回）
	AutioDurationProjectCount int `json:"auto_duration_project_count,omitempty"`
}
