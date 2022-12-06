package auth

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetPublicKeyRequest 获取公钥 API Request
type GetPublicKeyRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Credential 密钥对label，枚举值：“primary”或“backup”
	Credential enum.Credential `json:"credential,omitempty"`
}

// Encode implement GetRequest interface
func (r GetPublicKeyRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("credential", string(r.Credential))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetPublicKeyResponse 获取公钥 API Response
type GetPublicKeyResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// PubKey 公钥
		PubKey *PublicKey `json:"pubkey,omitempty"`
	} `json:"data,omitempty"`
}
