package keyword

import "encoding/json"

// DeleteRequest 删除关键词 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 待删除关键词的计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// KeywordIDs 待删除的关键词id列表，一次最多批量删除100个关键词
	KeywordIDs []uint64 `json:"keyword_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r DeleteRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
