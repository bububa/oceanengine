package oauth

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/oauth"
)

// RefreshToken 刷新Refresh Token
// 由于Access_Token有效期（默认1天）较短,当Access_Token超时后，可以使用refresh_token进行刷新，每次刷新都会产生新的access_token和Refresh_Token，同时重置二者的有效期。
// Refresh_Token有效期是30天，但是刷新后会产生新的Refresh_token，老的Refresh_token会过期
func RefreshToken(ctx context.Context, clt *core.SDKClient, refreshToken string) (*oauth.AccessTokenResponseData, error) {
	req := &oauth.AccessTokenRequest{
		AppId:        clt.AppID,
		Secret:       clt.Secret,
		GrantType:    "refresh_token",
		RefreshToken: refreshToken,
	}
	var resp oauth.AccessTokenResponse
	err := clt.Post(ctx, "oauth2/refresh_token/", req, &resp, "")
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
