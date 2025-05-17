package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// DetailGet 获取商品列表
func DetailGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.DetailGetRequest) (*dpa.DetailGetResponseData, error) {
	var resp dpa.DetailGetResponse
	if err := clt.Get(ctx, "2/dpa/detail/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
