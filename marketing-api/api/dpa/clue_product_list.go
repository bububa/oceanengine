package dpa

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// ClueProductList 获取升级版商品列表
func ClueProductList(clt *core.SDKClient, accessToken string, req *dpa.ClueProductListRequest) (*dpa.ClueProductListResult, error) {
	var resp dpa.ClueProductListResponse
	if err := clt.GetAPI("2/dpa/clue_product/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
