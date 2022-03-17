package ad

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateBidRequest 更新计划出价 API Request
type UpdateBidRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Data 更新计划出价的列表，最多支持10个
	Data []UpdateBidRequestData `json:"data,omitempty"`
}

// UpdateBidRequestData 更新预算的计划id和出价
type UpdateBidRequestData struct {
	// AdID 广告计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// Bid 计划更新之后的出价，最多只有两位小数
	Bid float64 `json:"bid,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateBidRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
