package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// AppGet 查询应用信息
// 查询应用信息 可通过请求参数advertiser_id广告主id 查询广告主下应用列表及应用详细信息
func AppGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.AppListRequest) (*appmanagement.AppListResponseData, error) {
	var resp appmanagement.AppListResponse
	if err := clt.Get(ctx, "2/tools/app_management/app/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
