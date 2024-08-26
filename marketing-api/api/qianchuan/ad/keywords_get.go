package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// KeywordsGet 获取计划的搜索关键词
func KeywordsGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.KeywordsGetRequest) (*ad.KeywordsGetResult, error) {
	var resp ad.KeywordsGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/ad/keywords/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
