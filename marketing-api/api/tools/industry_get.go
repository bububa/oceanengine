package tools

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// IndustryGet 获取行业列表，通过接口可以获取到一级行业、二级行业、三级行业列表，其中代理商创建广告主时使用的是二级行业，而在创建创意填写创意分类时使用的是三级行业，请注意区分。
func IndustryGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *tools.IndustryGetRequest) ([]tools.Industry, error) {
	var resp tools.IndustryGetResponse
	err := clt.Get(ctx, "2/tools/industry/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
