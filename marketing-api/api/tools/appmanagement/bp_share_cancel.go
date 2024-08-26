package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// BpShareCancel 取消应用共享关系
// 取消应用共享关系
// 当前共享类型支持PART（指定账户共享）、ALL（组织内某类账户所有共享）和COMPANY （公司主体内某类型所有账户共享）。
func BpShareCancel(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.BpShareRequest) (*appmanagement.BpShareData, error) {
	var resp appmanagement.BpShareResponse
	if err := clt.Post(ctx, "2/tools/app_management/bp_share/cancel/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
