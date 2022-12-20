package v3

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeepBidTypeGetRequest 获取可用深度优化方式体验版 API Request
type DeepBidTypeGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AssetID 资产id
	AssetID uint64 `json:"asset_id,omitempty"`
	// ExternalAction 优化目标
	ExternalAction enum.AdConvertType `json:"external_action,omitempty"`
	// DeepExternalAction 深度优化目标，当优化目标不等于AD_CONVERT_TYPE_PAY时必填
	DeepExternalAction enum.DeepExternalAction `json:"deep_external_action,omitempty"`
}

// Encode implement GetRequest interface
func (r DeepBidTypeGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.AssetID > 0 {
		values.Set("asset_id", strconv.FormatUint(r.AssetID, 10))
	}
	values.Set("external_action", string(r.ExternalAction))
	if r.ExternalAction != enum.AD_CONVERT_TYPE_PAY {
		values.Set("deep_external_action", string(r.DeepExternalAction))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// DeepBidTypeGetResponse 获取可用深度优化方式体验版 API Response
type DeepBidTypeGetResponse struct {
	model.BaseResponse
	// Data 返回数据
	Data struct {
		// DeepBidType 可用深度优化方式
		DeepBidType []enum.DeepBidType `json:"deep_bid_type,omitempty"`
	} `json:"data"`
}
