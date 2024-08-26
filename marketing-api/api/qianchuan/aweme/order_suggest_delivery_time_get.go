package aweme

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/aweme"
)

// OrderSuggestDeliveryTimeGet 获取建议延长时长
// 在追加随心推订单预算时，支持该订单的获取建议投放时长
func OrderSuggestDeliveryTimeGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *aweme.OrderSuggestDeliveryTimeGetRequest) (float64, error) {
	var resp aweme.OrderSuggestDeliveryTimeGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/aweme/order/suggest/delivery_time/get/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.SuggestDeliveryTime, nil
}
