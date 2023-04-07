package v3

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// StopRequest 关停一键起量 API Request
type StopRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionIDs 广告id
	PromotionIDs []uint64 `json:"promotion_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r StopRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
