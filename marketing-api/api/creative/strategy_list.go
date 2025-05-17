package creative

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/creative"
)

// StrategyList 获取模板（白盒策略）列表
// 获取白盒策略列表，支持搜索
func StrategyList(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.StrategyListRequest) (*creative.StrategyListData, error) {
	var resp creative.StrategyListResponse
	if err := clt.Get(ctx, "v2/creative/strategy/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
