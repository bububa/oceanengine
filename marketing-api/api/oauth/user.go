package oauth

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/oauth"
)

// UserInfo 获取授权User信息
// API授权是以User为纬度的，Access Token记录了授权User信息；通过此接口可以获取每一个Access Token对应的User信息，方便开发者区分以及管理对应授权关系
func UserInfo(ctx context.Context, clt *core.SDKClient, accessToken string) (*oauth.UserInfoResponseData, error) {
	var resp oauth.UserInfoResponse
	err := clt.Get(ctx, "2/user/info", nil, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
