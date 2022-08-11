package tools

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// SearchBidRatioGetRequest 获取快投推荐出价系数 API Request
type SearchBidRatioGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 广告ID，修改广告时需要传
	AdID uint64 `json:"ad_id,omitempty"`
}

// Encode implement GetRequest interface
func (r SearchBidRatioGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.AdID > 0 {
		values.Set("ad_id", strconv.FormatUint(r.AdID, 10))
	}
	return values.Encode()
}

// SearchBidRatioGetResponse 获取快投推荐出价系数 API Response
type SearchBidRatioGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// Ratio 推荐的快投出价系数
		Ratio float64 `json:"ratio,omitempty"`
	} `json:"data,omitempty"`
}
