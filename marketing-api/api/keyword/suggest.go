package keyword

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/keyword"
)

// Suggest 搜索快投关键词推荐
// 搜索快投关键词推荐接口用于获取系统推荐的“搜索快投关键词”。
// 搜索快投关键词仅适用于开启“搜索快投”功能的信息流广告计划。 本接口一次最多仅返回1000个关键词，且当前仅支持返回“广泛匹配”的关键词。
func Suggest(ctx context.Context, clt *core.SDKClient, accessToken string, req *keyword.SuggestRequest) ([]keyword.SuggestKeyword, error) {
	var resp keyword.SuggestResponse
	err := clt.Get(ctx, "2/keyword_feedads/suggest/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.List, nil
}
