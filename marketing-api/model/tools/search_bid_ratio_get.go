package tools

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
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
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.AdID > 0 {
		values.Set("ad_id", strconv.FormatUint(r.AdID, 10))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
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
