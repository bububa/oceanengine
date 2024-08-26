package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/ad"
)

// UpdateBudget 更新计划预算
// 通过此接口用于更新广告计划的预算；
// 一次可以处理100个计划
// 24小时内修改预算操作，不能超过20次，24小时是指自然天的24小时；
// 单次修改预算幅度不能低于100元（增加或者减少）；
// 修改后预算金额，不能低于当前已消费金额的105%，以整百单位向上取整；
func UpdateBudget(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.UpdateBudgetRequest) (*ad.UpdateResponseData, error) {
	var resp ad.UpdateResponse
	err := clt.Post(ctx, "2/ad/update/budget/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
