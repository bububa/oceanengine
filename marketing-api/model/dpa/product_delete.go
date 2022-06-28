package dpa

import (
	"encoding/json"
)

// ProductDeleteRequest
type ProductDeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PlatformID 商品库ID
	PlatformID uint64 `json:"platform_id,omitempty"`
	// ProductID 商品ID
	ProductID uint64 `json:"product_id,omitempty"`
}

// Encode implement PostRequest interface
func (r ProductDeleteRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}
