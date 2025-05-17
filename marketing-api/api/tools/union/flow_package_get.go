package union

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/union"
)

// FLowPackageGet 获取穿山甲流量包。
// “运营推荐”流量包是由客户对接的运营同学按照客户的需要，手动配置并推送给客户使用的，查询前请和对接的运营同学确认是否有运营推荐流量包，并明确定向/排除条件。
func FlowPackageGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *union.FlowPackageGetRequest) (*union.FlowPackageGetData, error) {
	var resp union.FlowPackageGetResponse
	if err := clt.Get(ctx, "2/tools/union/flow_package/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
