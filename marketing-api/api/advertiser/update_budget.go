package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// UpdateBudget 更新账户日预算
// 此接口可以更新广告主账号设置的预算类型与预算。
// 账户预算类别：不限预算、日预算
// 范围1000<=X <= 9999999.99、仅支持最多2位小数
func UpdateBudget(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.UpdateBudgetRequest) error {
	return clt.Post(ctx, "2/advertiser/update/budget/", req, nil, accessToken)
}
