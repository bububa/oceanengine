package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/tools/adraise/v3"
)

// SuggestBudgetGet 获取广告建议起量预算
func SuggestBudgetGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.SuggestBudgetGetRequest) ([]v3.SuggestBudget, error) {
	var resp v3.SuggestBudgetGetResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/suggest_budget/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
