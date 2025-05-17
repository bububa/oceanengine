package campaign

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/campaign"
)

// Update 修改广告组
// 此接口用于修改广告组信息。
// 推广目的一旦创建是不允许修改的
// 当选择日预算类型时，日预算不少于300元；
// 24小时内修改预算操作，不能超过20次，24小时是指自然天的24小时；
// 单次修改预算幅度不能低于100元（增加或者减少）；
// 修改后预算金额，不能低于当前已消费金额的105%，以整百单位向上取整；
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.UpdateRequest) (uint64, error) {
	var resp campaign.UpdateResponse
	err := clt.Post(ctx, "2/campaign/update/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.CampaignID, nil
}
