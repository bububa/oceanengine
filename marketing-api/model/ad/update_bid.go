package ad

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateBidRequest 更新计划出价 API Request
type UpdateBidRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Data 批量修改，包含计划id和出价
	Data []UpdateBidRequestList `json:"data,omitempty"`
}

// UpdateBidRequestList 修改信息
type UpdateBidRequestList struct {
	// AdID 广告计划ID，广告计划ID需要属于广告主
	AdID uint64 `json:"ad_id,omitempty"`
	// Bid 出价，单位“元”，精度：两位小数。
	Bid float64 `json:"bid,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateBidRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
