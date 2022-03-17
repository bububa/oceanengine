package landinggroup

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SiteOptStatusUpdateRequest 更新落地页组站点状态 API Request
type SiteOptStatusUpdateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// GroupID 落地页组 ID
	GroupID uint64                    `json:"group_id,omitempty"`
	Data    []SiteOptStatusUpdateItem `json:"data,omitempty"`
}

// SiteOptStatusUpdateItem 更新项
type SiteOptStatusUpdateItem struct {
	// MemberID 成员 ID，即站点在落地页组中的唯一标识
	MemberID uint64 `json:"member_id,omitempty"`
	// SiteOptStatus 站点启用状态
	SiteOptStatus enum.SiteOptStatus `json:"site_opt_status,omitempty"`
}

// Encode implement PostRequest interface
func (r SiteOptStatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
