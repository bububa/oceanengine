package promotion

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/promotion"
)

// BudgetUpdate 更新广告预算
func BudgetUpdate(clt *core.SDKClient, accessToken string, req *promotion.BudgetUpdateRequest) (*promotion.UpdateResponseData, error) {
	var resp promotion.UpdateResponse
	err := clt.Post("v3.0/promotion/budget/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
