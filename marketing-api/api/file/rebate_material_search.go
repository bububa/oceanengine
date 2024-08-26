package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// RebateMaterialSearch 【代理商】明点无效素材查询
// 支持根据代理商id和月份日期等入参数，查找返点政策粒度下的有效/无效素材
func RebateMaterialSearch(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.RebateMaterialSearchRequest) (*file.RebateMaterialSearchResult, error) {
	var resp file.RebateMaterialSearchResponse
	if err := clt.Get(ctx, "2/file/rebate/material_search/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
