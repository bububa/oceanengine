package promotion

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeleteRequest 批量删除广告 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionIDs 广告ID集合，list长度限制1～10
	PromotionIDs []uint64 `json:"promotion_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
