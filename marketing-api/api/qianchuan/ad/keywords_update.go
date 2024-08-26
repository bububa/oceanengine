package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// KeywordsUpdate 更新关键词
// 用于更新搜索计划关键词
func KeywordsUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.KeywordsUpdateRequest) error {
	return clt.Post(ctx, "v1.0/qianchuan/ad/keywords/update/", req, nil, accessToken)
}
