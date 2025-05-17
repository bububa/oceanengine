package wechat

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/wechat"
)

// BpAssetManagementShare 设置微信小游戏/小程序共享
// 通用的共享接口，目前支持微信小游戏、小程序资产
// 通过该接口，将微信小游戏/小程序资产共享给指定账户
func BpAssetManagementShare(ctx context.Context, clt *core.SDKClient, accessToken string, req *wechat.BpAssetManagementShareRequest) ([]wechat.BpAssetManagementShareError, error) {
	var resp wechat.BpAssetManagementShareResponse
	if err := clt.PostAPI(ctx, "v3.0/tools/bp_asset_management/share/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.ErrorList, nil
}
