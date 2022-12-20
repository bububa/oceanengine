package v3

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OptimizedGoalGetRequest 获取优化目标 API Request
type OptimizedGoalGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// LandingType 广告组推广目的，允许值:LINK 销售线索收集
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// AdType 广告类型，允许值： ALL
	AdType enum.CampaignType `json:"ad_type,omitempty"`
	// AssetType 资产类型，允许值:THIRD_EXTERNAL 三方落地页、TETRIS_EXTERNAL 建站、APP 应用、QUICK_APP 快应用、MINI_PROGRAME字节小程序
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// AssetID 三方的资产id，当asset_type为THIRD_EXTERNAL时必填
	AssetID uint64 `json:"asset_id,omitempty"`
	// PackageName 应用包名称
	PackageName string `json:"package_name,omitempty"`
	// AppType 应用类型，当asset_type为应用APP时必填
	// 可选值：ANDROID 、IOS
	AppType string `json:"app_type,omitempty"`
	// AppPromotionType 子目标，当 landing_type = APP 时该字段必填，允许值： DOWNLOAD 应用下载、LAUNCH 应用调用、RESERVE 预约下载
	AppPromotionType enum.AppPromotionType `json:"app_promotion_type,omitempty"`
}

// Encode implement GetRequest interface
func (r OptimizedGoalGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("landing_type", string(r.LandingType))
	values.Set("app_type", string(r.AppType))
	values.Set("asset_type", string(r.AssetType))
	if r.AssetID > 0 {
		values.Set("asset_id", strconv.FormatUint(r.AssetID, 10))
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
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// OptimizedGoalGetResponse 获取优化目标 API Response
type OptimizedGoalGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *OptimizedGoalGetResponseData `json:"data,omitempty"`
}

// OptimizedGoalGetResponseData json返回值
type OptimizedGoalGetResponseData struct {
	// AssetIDs 资产 id
	AssetIDs []uint64 `json:"asset_ids,omitempty"`
	// Goals 优化目标数据列表
	Goals []eventmanager.EventConvertOptimizedGoal `json:"goals,omitempty"`
}
