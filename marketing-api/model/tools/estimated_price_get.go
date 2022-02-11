package tools

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// EstimatedPriceGetRequest 获取预估点击成本 API Request
type EstimatedPriceGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r EstimatedPriceGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return values.Encode()
}

// EstimatedPriceGetResponse 获取预估点击成本 API Response
type EstimatedPriceGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *EstimatedPrice `json:"data,omitempty"`
}

// EstimatedPrice 预估点击成本
type EstimatedPrice struct {
	// UpperBound 建议出价上界，单位为元
	UpperBound float64 `json:"upper_bound,omitempty"`
	// LowerBound 建议出价下界，单位为元
	LowerBound float64 `json:"lower_bound,omitempty"`
	// CpcPrice 回填建议出价，单位为元
	CpcPrice float64 `json:"cpc_price,omitempty"`
}
