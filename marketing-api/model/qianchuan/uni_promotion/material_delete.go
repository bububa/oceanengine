package unipromotion

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MaterialDeleteRequest 删除全域推广计划下素材 API Request
type MaterialDeleteRequest struct {
	// AdvertiserID 千川广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// MaterialIDs 素材id list
	// 注意：最多支持100个素材
	MaterialIDs []uint64 `json:"material_ids,omitempty"`
	// Titles 需要删除的title list
	// 注意：完全匹配，最多支持100个title
	Titles []string `json:"titles,omitempty"`
}

// Encode implements PostRequest inteface
func (r MaterialDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
