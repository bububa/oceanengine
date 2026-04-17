package project

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/project"
)

// CPABidUpdate 批量更新项目出价
func CPABidUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *project.CPABidUpdateRequest) (*project.UpdateResponseData, error) {
	var resp project.UpdateResponse
	if err := clt.PostAPI(ctx, "v3.0/project/cpa_bid/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
