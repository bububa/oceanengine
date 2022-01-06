package comment

import (
	"encoding/json"
)

// TermsBandedUpdateRequest 更新屏蔽词 API Request
type TermsBandedUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// OriginTerms 待更新的屏蔽词; 屏蔽词长度范围为0-100
	OriginTerms []string `json:"origin_terms,omitempty"`
	// 更新后的屏蔽词; 如果new_terms已存在，则等同于删除origin_terms；如果origin_terms不存在，则等同于新增new_terms; 屏蔽词长度范围为0-100
	NewTerms []string `json:"new_terms,omitempty"`
}

// Encode implement PostRequest interface
func (r TermsBandedUpdateRequest) Encode() []byte {
	js, _ := json.Marshal(r)
	return js
}
