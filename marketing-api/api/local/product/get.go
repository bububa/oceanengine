package product

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/product"
)

// Get 获取可投商品列表
func Get(ctx context.Context, clt *core.SDKClient, accessToken string, req *product.GetRequest) (*product.GetResult, error) {
	var resp product.GetResponse
	if err := clt.GetAPI(ctx, "v3.0/local/product/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
