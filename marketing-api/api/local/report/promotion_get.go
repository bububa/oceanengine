package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/report"
)

// PromotionGet 获取广告数据
// 获取广告维度报表数据内容
func PromotionGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.PromotionGetRequest) (*report.PromotionGetResult, error) {
	var resp report.PromotionGetResponse
	if err := clt.GetAPI(ctx, "v3.0/local/report/promotion/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
