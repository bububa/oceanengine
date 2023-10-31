package oauth

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AppAccessTokenRequest 获取APP Access Token API Request
type AppAccessTokenRequest struct {
	// AppID 开发者申请的应用APP_ID，可通过“应用管理”界面查看
	AppID uint64 `json:"app_id,omitempty"`
	// Secret 开发者应用的私钥Secret，可通过“应用管理”界面查看（确保填入secret与app_id对应以免报错！）
	Secret string `json:"secret,omitempty"`
}

// Encode implement PostRequest interface
func (r AppAccessTokenRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AppAccessTokenResponse 获取App Access Token API Response
type AppAccessTokenResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AppAccessTokenResult `json:"data,omitempty"`
}

// AppAccessTokenResult
type AppAccessTokenResult struct {
	// AccessToken 应用级别token
	AccessToken string `json:"access_token,omitempty"`
	// ExpiresIn access_token剩余有效时间,单位(秒)
	ExpiresIn int64 `json:"expires_in,omitempty"`
}
