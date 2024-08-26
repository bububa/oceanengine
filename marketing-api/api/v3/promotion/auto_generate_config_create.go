package promotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/promotion"
)

// AutoGenerateConfigCreate 新建/修改白盒配置（广告升级版）
// 用来保存广告主自定义的策略及物料，用于创意智能生成。创建时promotion_id选填，更新时promotion_id必填，多次更新返回的configid不变
func AutoGenerateConfigCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *promotion.AutoGenerateConfigCreateRequest) (uint64, error) {
	var resp promotion.AutoGenerateConfigCreateResponse
	if err := clt.PostAPI(ctx, "v3.0/promotion/auto_generate_config/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.ConfigID, nil
}
