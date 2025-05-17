package duoplus

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/duoplus"
)

// OrderList 查询订单列表
// 查询DOU+订单属性，支持过滤
func OrderList(ctx context.Context, clt *core.SDKClient, accessToken string, req *duoplus.OrderListRequest) (*duoplus.OrderListResult, error) {
	var resp duoplus.OrderListResponse
	if err := clt.GetAPI(ctx, "v3.0/duoplus/order/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
