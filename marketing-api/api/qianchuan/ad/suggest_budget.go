package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// SuggestBudget 获取建议预算接口
func SuggestBudget(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.SuggestBudgetRequest) (*ad.SuggestBudgetResult, error) {
	var resp ad.SuggestBudgetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/suggest/budget/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
