package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// RecommendKeywordsGet 获取系统推荐的搜索关键词
// 获取巨量千川广告账户下，系统推荐的搜索关键词，以用于创编搜索广告
func RecommendKeywordsGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.RecommendKeywordsGetRequest) (*ad.RecommendKeywordsGetResult, error) {
	var resp ad.RecommendKeywordsGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/ad/recommend_keywords/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
