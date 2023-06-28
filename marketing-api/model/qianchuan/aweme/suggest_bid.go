package aweme

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SuggestBidRequest 获取随心推短视频建议出价 API Request
type SuggestBidRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// DeliverySetting 投放设置
	DeliverySetting *DeliverySetting `json:"delivery_setting,omitempty"`
	// Audience 定向设置
	Audience *Audience `json:"audience,omitempty"`
}

// Encode implement GetRequest interface
func (r SuggestBidRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.DeliverySetting != nil {
		values.Set("delivery_setting", string(util.JSONMarshal(r.DeliverySetting)))
	}
	if r.Audience != nil {
		values.Set("audience", string(util.JSONMarshal(r.Audience)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// SuggestBidResponse 获取随心推短视频建议出价 API Response
type SuggestBidResponse struct {
	model.BaseResponse
	Data struct {
		// SuggestBid 建议出价
		SuggestBid float64 `json:"suggest_bid,omitempty"`
	} `json:"data,omitempty"`
}
