package analyse

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/product/analyse"
)

// List 获取商品竞争分析列表
// 获取商品竞争分析列表
func List(clt *core.SDKClient, accessToken string, req *analyse.ListRequest) (*analyse.ListResponseData, error) {
	var resp analyse.ListResponse
	if err := clt.Get("v1.0/qianchuan/product/analyse/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
