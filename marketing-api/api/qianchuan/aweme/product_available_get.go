package aweme

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/aweme"
)

// ProductAvailableGet 达人获取可投商品列表
func ProductAvailableGet(clt *core.SDKClient, accessToken string, req *aweme.ProductAvailableGetRequest) (*aweme.ProductAvailableGetResult, error) {
	var resp aweme.ProductAvailableGetResponse
	if err := clt.Get("v1.0/qianchuan/aweme/product/available/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
