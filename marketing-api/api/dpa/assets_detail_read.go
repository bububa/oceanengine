package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// AssetsDetailRead 获取投放条件详情
func AssetsDetailRead(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.AssetsDetailReadRequest) ([]dpa.Asset, error) {
	var resp dpa.AssetsDetailReadResponse
	if err := clt.Get(ctx, "2/dpa/assets/detail/read/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
