package interestaction

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/interestaction"
)

// KeywordSuggest 获取行为兴趣推荐关键词
func KeywordSuggest(ctx context.Context, clt *core.SDKClient, accessToken string, req *interestaction.KeywordSuggestRequest) ([]interestaction.Object, error) {
	var resp interestaction.KeywordSuggestResponse
	err := clt.Get(ctx, "2/tools/interest_action/keyword/suggest/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.Keywords, nil
}
