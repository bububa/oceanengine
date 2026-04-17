package project

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/project"
)

// DeepCPABidUpdate 批量更新深层项目出价
func DeepCPABidUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *project.DeepCPABidUpdateRequest) (*project.UpdateResponseData, error) {
	var resp project.UpdateResponse
	if err := clt.PostAPI(ctx, "v3.0/project/deep_cpa_bid/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
