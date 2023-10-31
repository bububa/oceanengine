package oauth

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AccessTokenRequest 获取Access Token API Request
type AccessTokenRequest struct {
	// AppId 开发者申请的应用APP_ID
	AppId uint64 `json:"app_id,omitempty"`
	// Secret 开发者应用的私钥Secret
	Secret string `json:"secret,omitempty"`
	// GrantType 授权类型。允许值: "auth_code"
	GrantType string `json:"grant_type,omitempty"`
	// AuthCode 授权码，在授权完成后回调时会提供该授权码，只有10分钟有效期，且只能使用一次，获取详情可见OAuth2.0授权
	AuthCode string `json:"auth_code,omitempty"`
	// RefreshToken 刷新token,从"获取Access Token"和“刷新Access Token”的返回结果中得到)，刷新后会过期，请及时保存最新的token
	RefreshToken string `json:"refresh_token,omitempty"`
}

// Encode implement PostRequest interface
func (r AccessTokenRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AccessTokenResponse 获取Access Token API Response
type AccessTokenResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AccessTokenResponseData `json:"data,omitempty"`
}

// AccessTokenResponseData json返回值
type AccessTokenResponseData struct {
	// AccessToken 用于验证权限的token
	AccessToken string `json:"access_token,omitempty"`
	// RefreshToken 刷新access_token,用于获取新的access_token和refresh_token，并且刷新过期时间
	RefreshToken string `json:"refresh_token,omitempty"`
	// ExpiresIn access_token剩余有效时间,单位(秒)
	ExpiresIn int64 `json:"expires_in,omitempty"`
	// RefreshTokenExpiresIn refresh_token剩余有效时间,单位(秒)
	RefreshTokenExpiresIn int64 `json:"refresh_token_expires_in,omitempty"`
	// AdvertiserID （将废弃，当前支持代理商角色账号授权可一次授权多个账号，请使用advertiser_ids字段获取授权账号ID）登录用户对应的广告账户ID, 如果授权多个广告主默认为第一个
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdvertiserIDs 授权的账户id列表；该字段返回授权选择的账户id，包括管家、代理商、广告主等，具体角色说明可以通过【获取已授权账户】列表来获取
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
}
