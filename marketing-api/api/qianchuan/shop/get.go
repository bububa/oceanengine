package shop

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/shop"
)

// Get 获取店铺账户信息
func Get(clt *core.SDKClient, accessToken string, req *shop.GetRequest) ([]shop.Shop, error) {
	var resp shop.GetResponse
	err := clt.Get("v1.0/qianchuan/shop/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
