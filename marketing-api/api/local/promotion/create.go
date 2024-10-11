package promotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/promotion"
)

// Create 创建广告
// 广告创建成功后，默认为启用状态
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *promotion.CreateRequest) (uint64, error) {
	var resp promotion.CreateResponse
	if err := clt.PostAPI(ctx, "v3.0/local/promotion/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.PromotionID, nil
}
