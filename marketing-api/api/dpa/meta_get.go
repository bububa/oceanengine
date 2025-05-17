package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// MetaGet 获取商品库元信息
func MetaGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.MetaGetRequest) ([]dpa.Meta, error) {
	var resp dpa.MetaGetResponse
	err := clt.Get(ctx, "2/dpa/meta/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
