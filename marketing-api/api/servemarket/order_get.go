package servemarket

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/servemarket"
)

// OrderGet 获取应用订单数据
func OrderGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *servemarket.OrderGetRequest) (*servemarket.OrderGetResponseData, error) {
	var resp servemarket.OrderGetResponse
	if err := clt.OpenGet(ctx, "v1.0/serve_market/order/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
