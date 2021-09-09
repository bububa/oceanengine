package tools

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// AdQualityGet 查询广告质量度
// 查询计划的广告质量度，只有产生过投放消耗的计划才会有质量度数据。
func AdQualityGet(clt *core.SDKClient, accessToken string, req *tools.AdQualityGetRequest) ([]tools.AdQuality, error) {
	var resp tools.AdQualityGetResponse
	err := clt.Get("2/tools/ad_quality/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.List, nil
}
