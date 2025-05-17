package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager"
	v3 "github.com/bububa/oceanengine/marketing-api/model/eventmanager/v3"
)

// OptimizedGoalGet 获取优化目标
// 查询事件管理下资产的优化目标
func OptimizedGoalGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.OptimizedGoalGetRequest) (*eventmanager.EventConvertOptimizedGoalGetResult, error) {
	var resp eventmanager.EventConvertOptimizedGoalGetResponse
	if err := clt.GetAPI(ctx, "v3.0/event_manager/optimized_goal/get_v2/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
