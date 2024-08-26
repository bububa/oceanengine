package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// BudgetGet 获取账户日预算
// 此接口可以获取广告主账号设置的预算类型与预算，可以一次查询100个广告主账号预算；
func BudgetGet(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserIDs []uint64) ([]advertiser.BudgetGetResponseList, error) {
	req := &advertiser.BudgetGetRequest{
		AdvertiserIDs: advertiserIDs,
	}
	var resp advertiser.BudgetGetResponse
	err := clt.Get(ctx, "2/advertiser/budget/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.List, nil
}
