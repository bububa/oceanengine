package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/keyword"
	v3 "github.com/bububa/oceanengine/marketing-api/model/keyword/v3"
)

// Suggest 搜索快投关键词推荐
// 快投2.0获取推荐关键词
func Suggest(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.SuggestRequest) ([]keyword.SuggestKeyword, error) {
	var resp keyword.SuggestResponse
	if err := clt.PostAPI(ctx, "v3.0/sugg_words/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.List, nil
}
