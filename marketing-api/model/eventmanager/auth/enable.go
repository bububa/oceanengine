package auth

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// EnableRequest 开启鉴权 API Request
type EnableRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement PostRequest interface
func (r EnableRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
