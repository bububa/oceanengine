package wechat

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/wechat"
)

// BpAssetManagementShareCancel 取消微信小游戏/小程序共享关系
func BpAssetManagementShareCancel(ctx context.Context, clt *core.SDKClient, accessToken string, req *wechat.BpAssetManagementShareRequest) ([]wechat.BpAssetManagementShareError, error) {
	var resp wechat.BpAssetManagementShareResponse
	if err := clt.PostAPI(ctx, "v3.0/tools/bp_asset_management/share/cancel", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.ErrorList, nil
}
