package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// ProductCreate 创建DPA商品（无商品id）
func ProductCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.ProductCreateRequest) (uint64, error) {
	var resp dpa.ProductCreateResponse
	if err := clt.Post(ctx, "2/dpa/product/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.ProductID, nil
}
