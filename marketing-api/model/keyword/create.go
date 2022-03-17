package keyword

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建关键词 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 待创建搜索关键词的计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// Keywords 待添加搜索关键词列表，一次最多添加100个关键词，一个计划最多可添加1000
	Keywords []Keyword `json:"keywords,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
