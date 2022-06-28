package dpa

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// AssetsList 获取投放条件列表
func AssetsList(clt *core.SDKClient, accessToken string, req *dpa.AssetsListRequest) (*dpa.AssetsListResponseData, error) {
	var resp dpa.AssetsListResponse
	if err := clt.Get("2/dpa/assets/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
