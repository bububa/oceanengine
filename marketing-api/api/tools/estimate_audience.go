package tools

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// EstimateAudience 查询受众预估结果
// 通过此接口你可以查询计划的受众条件，预估其在今日头条、抖音短视频、火山小视频和西瓜视频中的覆盖用户量，当受众预估的结果低于3万时，建议重新设置受众选项。
func EstimateAudience(ctx context.Context, clt *core.SDKClient, accessToken string, req *tools.EstimateAudienceRequest) (*tools.EstimateAudienceResponseData, error) {
	var resp tools.EstimateAudienceResponse
	if err := clt.Get(ctx, "2/tools/estimate_audience/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
