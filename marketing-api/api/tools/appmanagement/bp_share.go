package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// BpShare 设置应用共享
// 设置应用共享，可通过该接口将应用共享给相关组织或指定账户
func BpShare(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.BpShareRequest) (*appmanagement.BpShareData, error) {
	var resp appmanagement.BpShareResponse
	if err := clt.Post(ctx, "2/tools/app_management/bp_share/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
