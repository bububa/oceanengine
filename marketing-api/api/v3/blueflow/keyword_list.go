package blueflow

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/blueflow"
)

// KeywordList 获取广告下可用蓝海关键词
func KeywordList(ctx context.Context, clt *core.SDKClient, accessToken string, req *blueflow.KeywordListRequest) ([]blueflow.Keyword, error) {
	var resp blueflow.KeywordListResponse
	if err := clt.GetAPI(ctx, "v3.0/blue_flow_keyword/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
