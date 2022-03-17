package coupon

import (
	"github.com/bububa/oceanengine/marketing-api/util"
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
	return util.JSONMarshal(r)
}
