package dpa

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// ClueProductDetail 获取升级版商品详情
func ClueProductDetail(clt *core.SDKClient, accessToken string, req *dpa.ClueProductDetailRequest) ([]dpa.Product, error) {
	var resp dpa.ClueProductDetailResponse
	if err := clt.GetAPI("2/dpa/clue_product/detail/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Products, nil
}
