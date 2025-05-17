package aweme

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/aweme"
)

// OrderBudgetAdd 追加随心推订单预算
// 可以使用该接口进行投放中订单的续费和续期操作
func OrderBudgetAdd(ctx context.Context, clt *core.SDKClient, accessToken string, req *aweme.OrderBudgetAddRequest) error {
	return clt.Post(ctx, "v1.0/qianchuan/aweme/order/budget/add/", req, nil, accessToken)
}
