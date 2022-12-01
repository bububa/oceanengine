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
	// RelatedProduct 关联产品投放相关
	RelatedProduct *RelatedProduct `json:"related_product,omitempty"`
	// DownloadURL 下载链接
	DownloadURL string `json:"download_url,omitempty"`
	// DownloadType 下载方式，枚举值：DOWNLOAD_URL 直接下载、EXTERNAL_URL 落地页下载
	DownloadType enum.DownloadType `json:"download_type,omitempty"`
	// DownloadMode 优先从系统应用商店下载（下载模式），枚举值：APP_STORE_DELIVERY 优先商店下载、 DEFAULT 默认下载
	DownloadMode enum.DownloadMode `json:"download_mode,omitempty"`
	// LaunchType 调起方式，枚举值： DIRECT_OPEN 直接调起、EXTERNAL_OPEN 落地页调起
	LaunchType enum.LaunchType `json:"launch_type,omitempty"`
	// OpenURL Deeplink直达链接
	OpenURL string `json:"open_url,omitempty"`
	// UlinkURL ulink直达链接
	UlinkURL string `json:"urlink_url,omitempty"`
	// SubscribeURL 预约下载链接
	SubscribeURL string `json:"subscribe_url,omitempty"`
	// AssetType 资产类型
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// MicroPromotionType 小游戏类型
	MicroPromotionType enum.MicroPromotionType `json:"micro_promotion_type,omitempty"`
	// OptimizeGoal 优化目标
	OptimizeGoal *OptimizeGoal `json:"optimize_goal,omitempty"`
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
