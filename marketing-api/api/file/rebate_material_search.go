package file

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// RebateMaterialSearch 获取低效素材List
func RebateMaterialSearch(clt *core.SDKClient, accessToken string, req *file.RebateMaterialSearchRequest) (*file.RebateMaterialSearchResult, error) {
	var resp file.RebateMaterialSearchResponse
	if err := clt.Get("2/file/rebate/material_search/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
