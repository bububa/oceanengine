package event

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/event"
)

// ConvertOptimizedGoalGet 获取优化目标
// 查询事件管理下资产的优化目标
func ConvertOptimizedGoalGet(clt *core.SDKClient, accessToken string, req *event.ConvertOptimizedGoalGetRequest) (*event.ConvertOptimizedGoalGetResponseData, error) {
	var resp event.ConvertOptimizedGoalGetResponse
	err := clt.Get("2/tools/event_convert/optimized_goal/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
