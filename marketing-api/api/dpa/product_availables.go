package dpa

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// ProductAvailables 获取商品库信息
// 仅支持查询广告主有权限访问的商品库 支持「商品推广」的推广目的包含：应用推广（APP）、销售线索收集（LINK）、电商店铺（SHOP）、商品推广（DPA） 其中仅商品推广（DPA）支持多商品推广，其余仅支持单商品推广
func ProductAvailables(clt *core.SDKClient, accessToken string, req *dpa.ProductAvailablesRequest) ([]dpa.Platform, error) {
	var resp dpa.ProductAvailablesResponse
	err := clt.Get("2/dpa/product/availables/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
