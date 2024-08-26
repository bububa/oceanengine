package oauth

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/oauth"
)

// AdvertiserGet 获取已授权账户
// 重新授权会覆盖前一次授权，需要确保此次授权已经勾选了想要的全部账户！！重新授权需要重新获取access token;
// 此接口可查询授权界面选择的角色账号id,包括广告主id、管家id、代理商id,具体可以通过“advertiser_role” 来区分；
// 如果是管家或者代理商角色，需要通过“获取管家账户列表”/“获取代理商账户列表”获取其下的广告主；
func AdvertiserGet(ctx context.Context, clt *core.SDKClient, accessToken string) ([]oauth.Advertiser, error) {
	req := &oauth.AdvertiserGetRequest{
		AppId:       clt.AppID,
		Secret:      clt.Secret,
		AccessToken: accessToken,
	}
	var resp oauth.AdvertiserGetResponse
	err := clt.Get(ctx, "oauth2/advertiser/get", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.List, nil
}
