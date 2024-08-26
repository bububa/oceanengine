package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/keyword"
	v3 "github.com/bububa/oceanengine/marketing-api/model/keyword/v3"
)

// List 获取关键词列表
// 根据过滤条件获取符合条件的所有关键词。
// 目前仅支持根据ad_id获取该计划下的关键词。
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.ListRequest) ([]keyword.Keyword, error) {
	var resp keyword.GetResponse
	if err := clt.GetAPI(ctx, "v3.0/keyword/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
