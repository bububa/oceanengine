package union

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/union"
)

// FlowPackageDelete 删除穿山甲流量包
func FlowPackageDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *union.FlowPackageDeleteRequest) (uint64, error) {
	var resp union.FlowPackageDeleteResponse
	if err := clt.Post(ctx, "2/tools/union/flow_package/delete/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.FlowPackageID, nil
}
