package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AddPublicKeyRequest 新增公钥 API Request
type AddPublicKeyRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Credential 密钥对label，枚举值：“primary”或“backup”
	Credential enum.Credential `json:"credential,omitempty"`
	// PubKey RSA公钥（1024位）文本内容，格式为X.509 Certificate Subject Public Key Info，可使用openssl生成，详见附录。注意：使用openssl生成的公钥是多行文本，需要使用换行符“\n”将多行文本转为单行。
	PubKey string `json:"pubkey,omitempty"`
	// Signature 使用RS256签名算法和私钥对字符串“BYTEDANCE”签名后的结果，可使用openssl计算，详见附录
	Signature string `json:"signature,omitepty"`
}

// NewPrivateKey 生成新私钥，返回pem file
func NewPrivateKey(privateKey *rsa.PrivateKey) (string, error) {
	if privKey, err := rsa.GenerateKey(rand.Reader, RSA_BITS); err != nil {
		return "", err
	} else {
		privateKey = privKey
	}
	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	if err := util.WritePrivateKeyPem(buf, privateKey); err != nil {
		return "", err
	}
	encoded := buf.String()
	return encoded, nil
}

// NewAddPublicKeyRequestWithPrivateKey 根据私钥生成新增公钥请求
func NewAddPublicKeyRequestWithPrivateKey(advertiserID uint64, credential enum.Credential, privateKey *rsa.PrivateKey) (*AddPublicKeyRequest, error) {
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return nil, err
	}
	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	//pem格式编码
	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	if err := pem.Encode(buf, &publicBlock); err != nil {
		return nil, err
	}
	sign, err := util.SignWithPrivateKey([]byte(RSA_SIGN_SRC), privateKey)
	if err := util.VerifyWithPublicKey([]byte(RSA_SIGN_SRC), sign, &publicKey); err != nil {
		return nil, err
	}
	return &AddPublicKeyRequest{
		AdvertiserID: advertiserID,
		Credential:   credential,
		PubKey:       buf.String(),
		Signature:    base64.StdEncoding.EncodeToString(sign),
	}, nil
}

// NewAddPublicKeyRequest 生成新增公钥请求
func NewAddPublicKeyRequest(advertiserID uint64, credential enum.Credential) (*AddPublicKeyRequest, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, RSA_BITS)
	if err != nil {
		return nil, err
	}
	return NewAddPublicKeyRequestWithPrivateKey(advertiserID, credential, privateKey)
}

// Encode implement PostRequest interface
func (r AddPublicKeyRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AddPublicKeyResponse 新增公钥 API Response
type AddPublicKeyResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// PubKey 公钥
		PubKey *PublicKey `json:"pubkey,omitempty"`
	} `json:"data,omitempty"`
}
