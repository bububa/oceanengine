package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// ProductDelete 删除DPA商品
func ProductDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.ProductDeleteRequest) error {
	return clt.Post(ctx, "2/dpa/product/delete/", req, nil, accessToken)
}
