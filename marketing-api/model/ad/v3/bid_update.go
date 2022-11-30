package v3

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// BidUpdateRequest 更新出价 API Request
type BidUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Data 批量修改，包含计划id和出价
	Data []BidUpdateData `json:"data,omitempty"`
}

// BidUpdateData 修改信息
type BidUpdateData struct {
	// PromotionID 广告计划ID，广告计划ID需要属于广告主
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// Bid 出价，单位“元”，精度：两位小数。
	Bid float64 `json:"deepbid,omitempty"`
}

// Encode implement PostRequest interface
func (r BidUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
