package v3

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/adraise/v3"
)

// SuggestBudgetGet 获取广告建议起量预算
func SuggestBudgetGet(clt *core.SDKClient, accessToken string, req *v3.SuggestBudgetGetRequest) ([]v3.SuggestBudget, error) {
	var resp v3.SuggestBudgetGetResponse
	if err := clt.Get("v3.0/tools/suggest_budget/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
