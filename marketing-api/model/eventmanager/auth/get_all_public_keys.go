package auth

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetAllPublicKeyRequest 获取全部公钥 API Request
type GetAllPublicKeyRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r GetAllPublicKeyRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetAllPublicKeyResponse 获取全部公钥 API Response
type GetAllPublicKeyResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// PubKey 公钥
		PubKeys []PublicKey `json:"pubkeys,omitempty"`
	} `json:"data,omitempty"`
}
