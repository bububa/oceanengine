package v3

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeleteRequest 删除关键词 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// KeywordIDs 待删除的关键词id列表，一次最多批量删除100个关键词
	KeywordIDs []uint64 `json:"keyword_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
