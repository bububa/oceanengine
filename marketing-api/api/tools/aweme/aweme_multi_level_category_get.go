package aweme

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/aweme"
)

// AwemeMultiLevelCategoryGet 查询抖音类目列表
// 查询抖音类目列表,获取抖音完整类目
func AwemeMultiLevelCategoryGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *aweme.AwemeMultiLevelCategoryGetRequest) ([]aweme.Category, error) {
	var resp aweme.AwemeMultiLevelCategoryGetResponse
	err := clt.Get(ctx, "2/tools/aweme_multi_level_category/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
