package aweme

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/aweme"
)

// AwemeCategoryTopAuthorGet 查询抖音类目下的推荐达人
// 查询抖音类目下的推荐达人，根据类目id查询抖音推荐达人。
func AwemeCategoryTopAuthorGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *aweme.AwemeCategoryTopAuthorGetRequest) ([]aweme.Author, error) {
	var resp aweme.AwemeCategoryTopAuthorGetResponse
	err := clt.Get(ctx, "2/tools/aweme_category_top_author/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
