package v3

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	ad "github.com/bububa/oceanengine/marketing-api/model/ad/v3"
)

// CostProtectStatusGet 批量获取计划成本保障状态
func CostProtectStatusGet(clt *core.SDKClient, accessToken string, req *ad.CostProtectStatusGetRequest) ([]ad.CostProtectStatus, error) {
	var resp ad.CostProtectStatusGetResponse
	if err := clt.Get("v3.0/promotion/cost_protect_status/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
