package adraise

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/adraise"
)

// Estimate 获取起量预估值: 用来获取起量预估值
func Estimate(clt *core.SDKClient, accessToken string, req *adraise.EstimateRequest) (int64, error) {
	var resp adraise.EstimateResponse
	err := clt.Get("2/tools/ad_raise_estimate/get/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.EstimateNum, nil
}
