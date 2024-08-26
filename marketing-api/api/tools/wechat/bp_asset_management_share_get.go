package wechat

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/wechat"
)

// BpAssetManagementShareGet 查看微信小游戏/小程序共享范围
func BpAssetManagementShareGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *wechat.BpAssetManagementShareGetRequest) (*wechat.BpAssetManagementShareList, error) {
	var resp wechat.BpAssetManagementShareGetResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/bp_asset_management/share/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
