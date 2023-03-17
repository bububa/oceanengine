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
	// DeliveryMode 投放模式，允许值：MANUAL手动投放(默认值）、PROCEDURAL自动投放 自动投放仅支持landing_type=APP或MICRO_GAME
	// 当marketing_goal= LIVE时，仅支持MANUAL手动投放
	DeliveryMode enum.DeliveryMode `json:"delivery_mode,omitempty"`
	// LandingType 推广目的，枚举值：APP 应用推广、LINK 销售线索推广、MICRO_GAME小游戏
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// AppPromotionType 子目标，枚举值：DOWNLOAD 应用下载、LAUNCH 应用调用、RESERVE 预约下载
	AppPromotionType enum.AppPromotionType `json:"app_promotion_type,omitempty"`
	// MarketingGoal 营销场景，枚举值：VIDEO_AND_IMAGE 短视频/图片
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// AdType 广告类型，枚举值：ALL
	AdType enum.CampaignType `json:"ad_type,omitempty"`
	// Name 项目名称
	Name string `json:"name,omitempty"`
	// SearchBidRatio 出价系数
	SearchBidRatio float64 `json:"search_bid_ratio,omitempty"`
	// AudienceExtend 定向拓展
	AudienceExtend string `json:"audience_extend,omitempty"`
	// Keywords 搜索关键词列表
	Keywords []Keyword `json:"keywords,omitempty"`
	// RelatedProduct 关联产品投放相关
	RelatedProduct *RelatedProduct `json:"related_product,omitempty"`
	// DownloadURL 下载链接
	DownloadURL string `json:"download_url,omitempty"`
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
	// 小程序类型，landing_type = MICRO_GAME 时有效且必填
	// 允许值： WECHAT_GAME 微信小游戏、WECHAT_APP微信小程序、BYTE_GAME字节小游戏、BYTE_APP字节小程序
	MicroPromotionType enum.MicroPromotionType `json:"micro_promotion_type,omitempty"`
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
	Data struct {
		// ProjectID 项目id
		ProjectID uint64 `json:"project_id,omitempty"`
	} `json:"data,omitempty"`
}
