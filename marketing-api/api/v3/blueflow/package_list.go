package blueflow

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/blueflow"
)

// PackageList 获取蓝海流量包
// 本接口支持查询搜索蓝海项目可用的蓝海流量包ID
// 您需要先通过本接口查询blue_flow_package_id，再前往「创建项目」接口创建搜索蓝海项目
func PackageList(ctx context.Context, clt *core.SDKClient, accessToken string, req *blueflow.PackageListRequest) ([]blueflow.Package, error) {
	var resp blueflow.PackageListResponse
	if err := clt.GetAPI(ctx, "v3.0/blue_flow_package/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.BlueFlowPackages, nil
}
