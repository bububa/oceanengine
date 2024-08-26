package union

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/union"
)

// FlowPackageCreate 创建穿山甲流量包
func FlowPackageCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *union.FlowPackageCreateRequest) (uint64, error) {
	var resp union.FlowPackageCreateResponse
	if err := clt.Post(ctx, "2/tools/union/flow_package/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.FlowPackageID, nil
}
