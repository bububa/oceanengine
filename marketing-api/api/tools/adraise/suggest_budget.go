package adraise

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/adraise"
)

// SuggestBudgetGet 获取广告建议起量预算
func SuggestBudgetGet(clt *core.SDKClient, accessToken string, req *adraise.SuggestBudgetGetRequest) ([]adraise.SuggestBudget, error) {
	var resp adraise.SuggestBudgetGetResponse
	if err := clt.Get("v3.0/tools/suggest_budget/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
