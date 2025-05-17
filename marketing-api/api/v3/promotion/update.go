package promotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/promotion"
)

// Update 修改广告
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *promotion.UpdateRequest) (*promotion.UpdateResponseData, error) {
	var resp promotion.UpdateResponse
	if err := clt.PostAPI(ctx, "v3.0/promotion/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
