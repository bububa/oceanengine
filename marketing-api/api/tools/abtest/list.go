package abtest

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/abtest"
)

// List 获取实验列表
// 该接口用于获取广告主账户下创建的AB实验，接口返回的实验列表默认按照创建时间降序。
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *abtest.ListRequest) (*abtest.ListData, error) {
	var resp abtest.ListResponse
	if err := clt.Get(ctx, "2/tools/ab_test_list/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
