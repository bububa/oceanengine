package aweme

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/aweme"
)

// OrderGet 获取随心推订单列表
// 此接口用于获取小店随心推订单列表
func OrderGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *aweme.OrderGetRequest) (*aweme.OrderGetResult, error) {
	var resp aweme.OrderGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/aweme/order/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
