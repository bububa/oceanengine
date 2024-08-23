package promotion

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/promotion"
)

// CostProtectStatusGet 批量获取计划成本保障状态
func CostProtectStatusGet(clt *core.SDKClient, accessToken string, req *promotion.CostProtectStatusGetRequest) ([]promotion.CostProtectStatus, error) {
	var resp promotion.CostProtectStatusGetResponse
	if err := clt.GetAPI("v3.0/promotion/cost_protect_status/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
