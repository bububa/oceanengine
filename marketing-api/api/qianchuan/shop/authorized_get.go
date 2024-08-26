package shop

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/shop"
)

// AuthorizedGet 获取广告主绑定的店铺列表
func AuthorizedGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *shop.AuthorizedGetRequest) (*shop.AuthorizedGetResult, error) {
	var resp shop.AuthorizedGetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/shop/authorized/get", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
