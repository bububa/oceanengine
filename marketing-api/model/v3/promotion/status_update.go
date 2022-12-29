package promotion

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// StatusUpdateRequest 更新广告状态 API Request
type StatusUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Data
	Data []StatusUpdateData `json:"data,omitempty"`
}

type StatusUpdateData struct {
	// PromotionID 广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// OptStatus  操作ENABLE启 用广告、DISABLE 暂停广告
	OptStatus enum.OptStatus `json:"opt_status,omitempty"`
}

// Encode implement PostRequest interface
func (r StatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
