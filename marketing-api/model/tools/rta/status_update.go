package rta

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// StatusUpdateRequest 批量启停账户下RTA策略 API Request
type StatusUpdateRequest struct {
	// AdvertiserID 广告账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// RtaIDs 预期设置的rta策略id
	RtaIDs []uint64 `json:"rta_ids,omitempty"`
	// Status 目标状态，允许值：启用 ENABLE、停用 DISABLE
	Status enum.RtaStatus `json:"status,omitempty"`
}

// Encode implement PostRequest interface
func (r StatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
