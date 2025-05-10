package promotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/promotion"
)

// MaterialStatusUpdate 批量更新广告素材启用状态
func MaterialStatusUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *promotion.MaterialStatusUpdateRequest) (*promotion.MaterialStatusUpdateResult, error) {
	var resp promotion.MaterialStatusUpdateResponse
	err := clt.PostAPI(ctx, "v3.0/material/status/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
