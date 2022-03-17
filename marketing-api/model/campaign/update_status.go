package campaign

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateStatusRequest 广告组更新状态 API Request
type UpdateStatusRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignIDs 广告组ID，不超过100个，且广告组ID属于广告主ID否则会报错；
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
	// OptStatus 操作, "enable"表示启用, "delete"表示删除, "disable"表示暂停；允许值: "enable", "delete", "disable",; 对于删除的广告组不可进行任何操作，否则会报错哈
	OptStatus string `json:"opt_status,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateStatusRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
