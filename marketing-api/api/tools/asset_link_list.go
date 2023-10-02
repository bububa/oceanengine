package tools

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// AssetLinkList 获取字节小程序/小游戏详情内容
// 获取字节小程序/小游戏详情内容
func AssetLinkList(clt *core.SDKClient, accessToken string, req *tools.AssetLinkListRequest) (*tools.AssetLinkListResult, error) {
	var resp tools.AssetLinkListResponse
	if err := clt.Get("v3.0/tools/asset_link/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
