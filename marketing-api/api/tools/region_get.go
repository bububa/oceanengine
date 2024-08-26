package tools

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// RegionGet 获取地域列表
// 获取地域列表（当前仅支持获取商圈ID）
func RegionGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *tools.RegionGetRequest) ([]tools.Region, error) {
	var resp tools.RegionGetResponse
	err := clt.Get(ctx, "2/tools/region/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
