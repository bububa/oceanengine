package promotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/promotion"
)

// List 获取广告列表
// 该接口暂不支持拉取 推广目的为获取线索的项目下的广告
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *promotion.ListRequest) (*promotion.ListResult, error) {
	var resp promotion.ListResponse
	if err := clt.GetAPI(ctx, "v3.0/local/promotion/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
