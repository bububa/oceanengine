package promotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/promotion"
)

// Update 更新广告
// 本接口为全量更新，每次更新会对所有参数进行调整，若希望进行局部更新，请先通过获取广告列表接口前置获取广告的详情信息再进行更新
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *promotion.UpdateRequest) error {
	return clt.PostAPI(ctx, "v3.0/local/promotion/update/", req, nil, accessToken)
}
