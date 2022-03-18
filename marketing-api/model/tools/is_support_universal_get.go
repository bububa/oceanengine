package tools

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// IsSupportUniversalGetRequest 查询是否支持通投智选 API Request
type IsSupportUniversalGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// LandingType 广告组推广目的，详见【附录-推广目的类型】，允许值：
	// APP 应用推广、ARTICLE 头条文章推广、AWEME 抖音号推广、DPA 商品目录推广、GOODS 商品推广、LINK 销售线索收集、QUICK_APP 快应用、SHOP 电商店铺推广
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// ExternalAction 优化目标，可通过【获取优化目标】接口获取
	ExternalAction enum.AdConvertType `json:"external_action,omitempty"`
	// DeepExternalAction 深度转化目标，可通过【获取优化目标】接口获取
	DeepExternalAction enum.DeepExternalAction `json:"deep_external_action,omitempty"`
	// DeepBidType 深度优化方式，允许值详见【附录-深度优化方式】
	DeepBidType enum.DeepBidType `json:"deep_bid_type,omitempty"`
}

// IsSupportUniversalGetResponse 查询是否支持通投智选 API Response
type IsSupportUniversalGetResponse struct {
	model.BaseResponse
	Data struct {
		IsSupportUniversal string `json:"is_support_universal,omitempty"`
	} `json:"data,omitempty"`
}
