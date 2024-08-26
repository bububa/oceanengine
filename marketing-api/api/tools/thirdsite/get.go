package thirdsite

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/thirdsite"
)

// Get 获取第三方落地页站点列表
// 通过此接口，广告主可以获取广告主下拥有的第三方落地页站点列表。
// 列表信息包含站点审核状态、站点创建时间、站点名称、站点ID、缩略图地址、站点地址。
func Get(ctx context.Context, clt *core.SDKClient, accessToken string, req *thirdsite.GetRequest) (*thirdsite.GetResponseData, error) {
	var resp thirdsite.GetResponse
	err := clt.Get(ctx, "2/tools/third_site/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
