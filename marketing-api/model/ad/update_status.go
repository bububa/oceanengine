package ad

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateStatusRequest 更新计划状态 API Request
type UpdateStatusRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdIDs 计划ID集合，限制1～100. 广告计划id需属于广告主
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// OptStatus  操作, "enable"表示启用, "delete"表示删除, "disable"表示暂停；允许值: "enable", "delete", "disable",; 对于删除的广告组不可进行任何操作，否则会报错哈
	OptStatus string `json:"opt_status,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateStatusRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
