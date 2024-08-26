package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// AssetsList 获取投放条件列表
func AssetsList(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.AssetsListRequest) (*dpa.AssetsListResponseData, error) {
	var resp dpa.AssetsListResponse
	if err := clt.Get(ctx, "2/dpa/assets/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
