package auth

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DisableRequest 关闭鉴权 API Request
type DisableRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement PostRequest interface
func (r DisableRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
