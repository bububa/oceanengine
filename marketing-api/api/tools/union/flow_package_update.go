package union

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/union"
)

// FlowPackageUpdate 修改穿山甲流量包
func FlowPackageUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *union.FlowPackageUpdateRequest) (uint64, error) {
	var resp union.FlowPackageUpdateResponse
	if err := clt.Post(ctx, "2/tools/union/flow_package/update/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.FlowPackageID, nil
}
