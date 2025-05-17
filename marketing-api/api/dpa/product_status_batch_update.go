package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// ProductStatusBatchUpdate 批量修改DPA商品状态
func ProductStatusBatchUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.ProductStatusBatchUpdateRequest) (*dpa.ProductStatusBatchUpdateResponseData, error) {
	var resp dpa.ProductStatusBatchUpdateResponse
	if err := clt.Post(ctx, "2/dpa/product_status/batch_update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
