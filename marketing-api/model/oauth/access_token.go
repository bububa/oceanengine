package oauth

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

type AccessTokenRequest struct {
	AppId        string `json:"app_id,omitempty"`
	Secret       string `json:"secret,omitempty"`
	GrantType    string `json:"grant_type,omitempty"`
	AuthCode     string `json:"auth_code,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (r AccessTokenRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

type AccessTokenResponse struct {
	model.BaseResponse
	Data *AccessTokenResponseData `json:"data,omitempty"`
}

type AccessTokenResponseData struct {
	AccessToken           string   `json:"access_token,omitempty"`
	RefreshToken          string   `json:"refresh_token,omitempty"`
	ExpiresIn             int64    `json:"expires_in,omitempty"`
	RefreshTokenExpiresIn int64    `json:"refresh_token_expires_in,omitempty"`
	AdvertiserID          uint64   `json:"advertiser_id,omitempty"`
	AdvertiserIDs         []uint64 `json:"advertiser_ids,omitempty"`
}
