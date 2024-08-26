package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// CategoryGet 获取DPA分类
func CategoryGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.CategoryGetRequest) ([]dpa.Category, error) {
	var resp dpa.CategoryGetResponse
	if err := clt.Get(ctx, "2/dpa/category/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
