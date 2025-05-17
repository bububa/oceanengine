package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// AssetsListV2 获取投放条件列表
func AssetsListV2(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.AssetsListV2Request) (*dpa.AssetsListV2Result, error) {
	var resp dpa.AssetsListV2Response
	if err := clt.GetAPI(ctx, "2/dpa/assets_v2/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
