package promotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/promotion"
)

// AutoGenerateConfigGet 查询配置详情
func AutoGenerateConfigGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *promotion.AutoGenerateConfigGetRequest) (*promotion.AutoGenerateConfig, error) {
	var resp promotion.AutoGenerateConfigGetResponse
	if err := clt.GetAPI(ctx, "v3.0/promotion/auto_generate_config/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
