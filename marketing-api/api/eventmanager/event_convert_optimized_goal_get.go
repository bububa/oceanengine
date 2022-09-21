package eventmanager

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager"
)

// EventConvertOptimizedGoalGet 获取优化目标
// 查询事件管理下资产的优化目标
func EventConvertOptimizedGoalGet(clt *core.SDKClient, accessToken string, req *eventmanager.EventConvertOptimizedGoalGetRequest) (*eventmanager.EventConvertOptimizedGoalGetResponseData, error) {
	var resp eventmanager.EventConvertOptimizedGoalGetResponse
	if err := clt.Get("2/tools/event_convert/optimized_goal/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
