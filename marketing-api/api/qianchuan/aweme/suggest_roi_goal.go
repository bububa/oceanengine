package aweme

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/aweme"
)

// SuggestRoiGoal 获取随心推ROI建议出价
// 获取随心推ROI建议出价
func SuggestRoiGoal(ctx context.Context, clt *core.SDKClient, accessToken string, req *aweme.SuggestRoiGoalRequest) (float64, error) {
	var resp aweme.SuggestRoiGoalResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/aweme/suggest/roi/goal/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.EcpRoiGoal, nil
}
