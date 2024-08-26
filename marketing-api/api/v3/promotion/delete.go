package promotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/promotion"
)

// Delete 批量删除广告
func Delete(ctx context.Context, clt *core.SDKClient, accessToken string, req *promotion.DeleteRequest) (*promotion.UpdateResponseData, error) {
	var resp promotion.UpdateResponse
	err := clt.PostAPI(ctx, "v3.0/promotion/delete/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
