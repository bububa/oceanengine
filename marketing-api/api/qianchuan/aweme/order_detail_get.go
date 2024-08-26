package aweme

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/aweme"
)

// OrderDetailGet 获取随心推订单详情
// 此接口用户获取小店随心推订单详情
func OrderDetailGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *aweme.OrderDetailGetRequest) (*aweme.Order, error) {
	var resp aweme.OrderDetailGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/aweme/order/detail/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
