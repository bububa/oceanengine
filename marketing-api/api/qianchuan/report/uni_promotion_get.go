package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// UniPromotionGet 全域推广数据
func UniPromotionGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.UniPromotionGetRequest) (*report.UniPromotionStats, error) {
	var resp report.UniPromotionGetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/report/uni_promotion/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
