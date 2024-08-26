package oauth

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/oauth"
)

// AppAccessToken 获取APP Access Token
func AppAccessToken(ctx context.Context, clt *core.SDKClient) (*oauth.AppAccessTokenResult, error) {
	req := &oauth.AppAccessTokenRequest{
		AppID:  clt.AppID,
		Secret: clt.Secret,
	}
	var resp oauth.AppAccessTokenResponse
	if err := clt.Post(ctx, "oauth2/app_access_token/", req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
