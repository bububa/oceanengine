package auth

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DelPublicKeyRequest 删除公钥 API Request
type DelPublicKeyRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Credential 密钥对label，枚举值：“primary”或“backup”
	Credential enum.Credential `json:"credential,omitempty"`
}

// Encode implement PostRequest interface
func (r DelPublicKeyRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
