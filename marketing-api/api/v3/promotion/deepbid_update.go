package promotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/promotion"
)

// DeepBidUpdate 更新深度出价
func DeepBidUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *promotion.DeepBidUpdateRequest) (*promotion.UpdateResponseData, error) {
	var resp promotion.UpdateResponse
	err := clt.PostAPI(ctx, "v3.0/promotion/deepbid/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
