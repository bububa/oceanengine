package dpa

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// ProductStatusBatchUpdate 批量修改DPA商品状态
func ProductStatusBatchUpdate(clt *core.SDKClient, accessToken string, req *dpa.ProductStatusBatchUpdateRequest) (*dpa.ProductStatusBatchUpdateResponseData, error) {
	var resp dpa.ProductStatusBatchUpdateResponse
	if err := clt.Post("2/dpa/product_status/batch_update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
