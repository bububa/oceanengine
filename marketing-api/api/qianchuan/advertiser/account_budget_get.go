package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/advertiser"
)

// AccountBudgetGet 获取账户日预算
func AccountBudgetGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.AccountBudgetGetRequest) (*advertiser.Budget, error) {
	var resp advertiser.AccountBudgetGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/account/budget/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
