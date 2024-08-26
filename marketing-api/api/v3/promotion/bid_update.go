package promotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/promotion"
)

// BidUpdate 更新出价
func BidUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *promotion.BidUpdateRequest) (*promotion.UpdateResponseData, error) {
	var resp promotion.UpdateResponse
	err := clt.PostAPI(ctx, "v3.0/promotion/bid/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
