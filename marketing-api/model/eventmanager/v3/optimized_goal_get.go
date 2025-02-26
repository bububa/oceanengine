package v3

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OptimizedGoalGetRequest 获取优化目标(巨量广告升级版) API Request
type OptimizedGoalGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// LandingType 广告组推广目的，允许值:LINK 销售线索收集
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// AdType 广告类型，允许值：ALL 可选值:
	// ALL 信息流
	// SEARCH 搜索
	AdType enum.CampaignType `json:"ad_type,omitempty"`
	// AssetType 资产类型，允许值:THIRD_EXTERNAL 三方落地页、TETRIS_EXTERNAL 建站、APP 应用、QUICK_APP 快应用、MINI_PROGRAME字节小程序
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// MultiAssetType 多投放载体，允许值：ORANGE_AND_AWEME 建站落地页和抖音主页
	MultiAssetType enum.MultiDeliveryMedium `json:"multi_asset_type,omitempty"`
	// SiteID 建站site_id，当asset_type为TETRIS_EXTERNAL时必填，site_id可以通过【获取橙子建站站点列表】接口获得
	SiteID uint64 `json:"site_id,omitempty"`
	// AssetID 三方的资产id，当asset_type为THIRD_EXTERNAL时必填
	AssetID uint64 `json:"asset_id,omitempty"`
	// QuickAppID 快应用id
	QuickAppID uint64 `json:"quick_app_id,omitempty"`
	// MiniProgramID 字节小程序资产id
	MiniProgramID string `json:"mini_program_id,omitempty"`
	// PackageName 应用包名称
	PackageName string `json:"package_name,omitempty"`
	// AppType 应用类型，当asset_type为应用APP时必填
	// 可选值：ANDROID 、IOS
	AppType string `json:"app_type,omitempty"`
	// AppPromotionType 子目标，可选值:
	// DOWNLOAD 应用下载
	// LAUNCH 应用调起
	// RESERVE 预约下载
	AppPromotionType enum.AppPromotionType `json:"app_promotion_type,omitempty"`
	// MarketingGoal 营销场景， 可选值:
	// LIVE
	// VIDEO_AND_IMAGE
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// DeliveryMode 投放模式 可选值:
	// MANUAL 手动投放模式
	// PROCEDURAL 自动投放模式
	DeliveryMode enum.DeliveryMode `json:"delivery_mode,omitempty"`
	// DpaAdType dpa广告类型，可选值:
	// DPA_APP 应用下载
	// DPA_LINK 落地页
	DpaAdType enum.DpaAdType `json:"dpa_adtype,omitempty"`
	// MicroPromotionType 小程序类型，landing_type = MICRO_GAME 时有效且必填
	// 可选值:
	// BYTE_APP 节小程序
	// BYTE_GAME 字节小游戏
	// WECHAT_APP 微信小程序
	// WECHAT_GAME 微信小游戏
	MicroPromotionType enum.MicroPromotionType `json:"micro_promotion_type,omitempty"`
	// MicroAppInstanceID 小程序资产id
	MicroAppInstanceID string `json:"micro_app_instance_id,omitempty"`
	// DeliveryType 可选值:
	// DURATION 周期稳投
	// NORMAL 常规投放
	DeliveryType enum.DeliveryType `json:"delivery_type,omitempty"`
}

// Encode implement GetRequest interface
func (r OptimizedGoalGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("landing_type", string(r.LandingType))
	values.Set("ad_type", string(r.AdType))
	values.Set("asset_type", string(r.AssetType))
	if r.SiteID > 0 {
		values.Set("site_id", strconv.FormatUint(r.SiteID, 10))
	}
	if r.AssetID > 0 {
		values.Set("asset_id", strconv.FormatUint(r.AssetID, 10))
	}
	if r.MiniProgramID != "" {
		values.Set("mini_program_id", r.MiniProgramID)
	}
	if r.PackageName != "" {
		values.Set("package_name", r.PackageName)
	}
	if r.AppType != "" {
		values.Set("app_type", r.AppType)
	}
	if r.AppPromotionType != "" {
		values.Set("app_promotion_type", string(r.AppPromotionType))
	}
	if r.MarketingGoal != "" {
		values.Set("marketing_goal", string(r.MarketingGoal))
	}
	if r.DeliveryMode != "" {
		values.Set("delivery_mode", string(r.DeliveryMode))
	}
	if r.DpaAdType != "" {
		values.Set("dpa_adtype", string(r.DpaAdType))
	}
	if r.MicroPromotionType != "" {
		values.Set("micro_promotion_type", string(r.MicroPromotionType))
	}
	if r.MicroAppInstanceID != "" {
		values.Set("micro_app_instance_id", r.MicroAppInstanceID)
	}
	if r.DeliveryType != "" {
		values.Set("delivery_type", string(r.DeliveryType))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}
