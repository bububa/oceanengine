package tools

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/tools"
)

// EstimateAudience 获取定向受众预估
// 定向受众预估，用于当选择定向后预估大概的受众数量以及曝光量
// 定向中的限制与创建计划中的限制保持一致
func EstimateAudience(ctx context.Context, clt *core.SDKClient, accessToken string, req *tools.EstimateAudienceRequest) (*tools.EstimateAudienceResult, error) {
	var resp tools.EstimateAudienceResponse
	if err := clt.GetAPI(ctx, "v1.0/qianchuan/tools/estimate_audience/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
