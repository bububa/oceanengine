package eventmanager

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// EventConvertOptimizedGoalGetRequest 获取优化目标 API Request
type EventConvertOptimizedGoalGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// LandingType 广告组推广目的，允许值:LINK 销售线索收集
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// MarketingPurpose 营销目的，允许值:CONVERSION 行动转化、INTENTION 用户意向、ACKNOWLEDGE 品牌认知
	MarketingPurpose enum.MarketingPurpose `json:"marketing_purpose,omitempty"`
	// AssetType 资产类型，允许值:THIRD_EXTERNAL 三方落地页、TETRIS_EXTERNAL 建站、APP 应用、QUICK_APP 快应用、MINI_PROGRAME字节小程序
	AssetType enum.AssetType `json:"asset_type,omitempty"`
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
	// CampaignType 广告类型，默认值信息流。允许值：FEED 信息流、SEARCH 搜索
	CampaignType enum.CampaignType `json:"campaign_type,omitempty"`
	// WechatAppID 微信小游戏原始id。获取方式：资产->青鸟线索通->转化组件->微信小游戏->小游戏列表，可获取小游戏原始id。
	WechatAppID string `json:"wechat_app_id,omitempty"`
}

// Encode implement GetRequest interface
func (r EventConvertOptimizedGoalGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("landing_type", string(r.LandingType))
	values.Set("marketing_purpose", string(r.MarketingPurpose))
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
	if r.CampaignType != "" {
		values.Set("campaign_type", string(r.CampaignType))
	}
	if r.WechatAppID != "" {
		values.Set("wechat_app_id", r.WechatAppID)
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// EventConvertOptimizedGoalGetResponse 获取优化目标 API Response
type EventConvertOptimizedGoalGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *EventConvertOptimizedGoalGetResult `json:"data,omitempty"`
}

// EventConvertOptimizedGoalGetResult json返回值
type EventConvertOptimizedGoalGetResult struct {
	// AssetIDs 资产 id
	AssetIDs []uint64 `json:"asset_ids,omitempty"`
	// Goals 优化目标数据列表
	Goals []EventConvertOptimizedGoal `json:"goals,omitempty"`
}
