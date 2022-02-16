package coupon

import (
	"encoding/json"
)

// EmployeeDeleteRequest 删除核销员 API Request
type EmployeeDeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// List 核销员列表
	List []Employee `json:"list,omitempty"`
}

// Encode implement PostRequest interface
func (r EmployeeDeleteRequest) Encode() []byte {
	b, _ := json.Marshal(r)
	return b
}
