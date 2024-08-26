package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// AndroidAppList 查询安卓应用信息（支持所有账户体系）
// 查询账户下安卓应用信息（支持所有账户体系）及应用详细信息
func AndroidAppList(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.AndroidAppListRequest) (*appmanagement.AppListResponseData, error) {
	var resp appmanagement.AppListResponse
	if err := clt.Get(ctx, "2/tools/app_management/android_app/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
