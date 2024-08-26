package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/advertiser"
)

// AccountBudgetUpdate 更新账户日预算
// 更新账户日预算，一天最多修改20次。
func AccountBudgetUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.AccountBudgetUpdateRequest) error {
	return clt.Post(ctx, "v1.0/qianchuan/account/budget/update/", req, nil, accessToken)
}
