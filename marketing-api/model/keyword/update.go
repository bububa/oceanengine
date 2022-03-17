package keyword

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateRequest 更新关键词 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 待创建搜索关键词的计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// Keywords 待更新搜索关键词列表，一次最多批量更新100个关键词
	Keywords []Keyword `json:"keywords,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
