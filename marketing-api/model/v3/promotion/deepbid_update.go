package promotion

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeepBidUpdateRequest 更新深度出价 API Request
type DeepBidUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Data 批量修改，包含计划id和出价
	Data []DeepBidUpdateData `json:"data,omitempty"`
}

// DeepBidUpdateData 修改信息
type DeepBidUpdateData struct {
	// PromotionID 广告计划ID，广告计划ID需要属于广告主
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// DeepBid 出价，单位“元”，精度：两位小数。
	DeepBid float64 `json:"deepbid,omitempty"`
}

// Encode implement PostRequest interface
func (r DeepBidUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
