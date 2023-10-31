package v3

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// BidSuggestRequest 查询建议出价（巨量广告升级版） API Request
type BidSuggestRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// Pricing 计费方式，可选值:
	// PRICING_CPC 按点击计费(CPC)
	// PRICING_CPM 按千次展示(CPM)
	// PRICING_OCPC 按转化优化,按点击计费
	// PRICING_OCPM 按转化优化,按展示计费
	Pricing enum.PricingType `json:"pricing,omitempty"`
	// ExternalAction 转化目标，允许值见【附录-枚举值-转化类型】
	ExternalAction enum.AdConvertType `json:"external_action,omitempty"`
	// DeepExternalAction 深度转化目标，允许值见【附录-枚举值-转化类型】
	DeepExternalAction enum.DeepExternalAction `json:"deep_external_action,omitempty"`
	// DeepBidType 深度转化目标，允许值见【附录-枚举值-深度优化方式】
	DeepBidType enum.DeepBidType `json:"deep_bid_type,omitempty"`
}

// Encode implement GetRequest interface
func (r BidSuggestRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("project_id", strconv.FormatUint(r.ProjectID, 10))
	values.Set("pricing", string(r.Pricing))
	values.Set("external_action", string(r.ExternalAction))
	if r.DeepExternalAction != "" {
		values.Set("deep_external_action", string(r.DeepExternalAction))
	}
	if r.DeepBidType != "" {
		values.Set("deep_bid_type", string(r.DeepBidType))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// BidSuggestResponse 查询建议出价（巨量广告升级版） API Response
type BidSuggestResponse struct {
	model.BaseResponse
	Data struct {
		Data *tools.BidSuggest `json:"data,omitempty"`
	} `json:"data,omitempty"`
}
