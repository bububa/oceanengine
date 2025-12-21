package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// HarmonyAppList 查询鸿蒙应用列表
func HarmonyAppList(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.HarmonyAppListRequest) (*appmanagement.AppListResponseData, error) {
	var resp appmanagement.AppListResponse
	if err := clt.Get(ctx, "2/tools/app_management/harmony_app/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
