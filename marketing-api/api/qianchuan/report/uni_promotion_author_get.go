package report

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// UniPromotionAuthorGet 获取全域推广抖音号维度数据
func UniPromotionAuthorGet(clt *core.SDKClient, accessToken string, req *report.UniPromotionDimensionGetRequest) (*report.UniPromotionDimensionGetResult, error) {
	var resp report.UniPromotionDimensionGetResponse
	err := clt.Get("v1.0/qianchuan/report/uni_promotion/dimension_data/author/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
