package promotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/promotion"
)

// Detail 获取广告详情
// 暂不支持拉取推广目的为获取线索的广告
func Detail(ctx context.Context, clt *core.SDKClient, accessToken string, req *promotion.DetailRequest) (*promotion.PromotionDetail, error) {
	var resp promotion.DetailResponse
	if err := clt.GetAPI(ctx, "v3.0/local/promotion/detail/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
