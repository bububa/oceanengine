package auth

import "github.com/bububa/oceanengine/marketing-api/enum"

// PublicKey 公钥
type PublicKey struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Credential 密钥对label，枚举值：“primary”或“backup”
	Credential enum.Credential `json:"credential,omitempty"`
	// KeyID
	KeyID uint64 `json:"key_id,omitempty"`
	// PubKey RSA公钥（1024位）文本内容，格式为X.509 Certificate Subject Public Key Info，可使用openssl生成，详见附录。注意：使用openssl生成的公钥是多行文本，需要使用换行符“\n”将多行文本转为单行。
	PubKey string `json:"pubkey,omitempty"`
}
