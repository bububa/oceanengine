package unipromotion

import "github.com/bububa/oceanengine/marketing-api/util"

// AuthInitRequest 全域授权初始化 API Request
type AuthInitRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implements PostRequest interface
func (r AuthInitRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
