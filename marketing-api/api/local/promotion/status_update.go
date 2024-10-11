package promotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/promotion"
)

// StatusUpdate 批量更新广告状态
func StatusUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *promotion.StatusUpdateRequest) (*promotion.UpdateResult, error) {
	var resp promotion.StatusUpdateResponse
	if err := clt.PostAPI(ctx, "v3.0/local/promotion/status/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
