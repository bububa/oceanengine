package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// ProductDetailGet 获取商品详情
func ProductDetailGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.ProductDetailGetRequest) (*dpa.DetailGetResponseData, error) {
	var resp dpa.DetailGetResponse
	if err := clt.Get(ctx, "2/dpa/product/detail/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
