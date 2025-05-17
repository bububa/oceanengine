package aweme

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/aweme"
)

// OrderCreate 创建随心推订单
func OrderCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *aweme.OrderCreateRequest) (*aweme.Order, error) {
	var resp aweme.OrderCreateResponse
	if err := clt.Post(ctx, "v1.0/qianchuan/aweme/order/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
