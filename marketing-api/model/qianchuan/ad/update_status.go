package ad

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
)

// UpdateStatusRequest 更新计划状态 API Request
type UpdateStatusRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdIDs 需要更新的广告计划id，最多支持10个
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// OptStatus 批量更新的广告计划状态
	OptStatus qianchuan.AdOptStatus `json:"opt_status,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateStatusRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
