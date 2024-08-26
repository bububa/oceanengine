package campaign

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/campaign"
)

// Create 创建广告组
// 此接口用于创建信息流广告组，对于搜索广告的创建可参照【搜索广告投放】
// 每个广告主账号下最多可允许创建500个广告组，如超出需要先删除一部分广告组后才可继续创建；
// 当选择日预算类型时，日预算不少于300元；
// 24小时内修改预算操作，不能超过20次，24小时是指自然天的24小时；
// 单次修改预算幅度不能低于100元（增加或者减少）；
// 修改后预算金额，不能低于当前已消费金额的105%，以整百单位向上取整；
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.CreateRequest) (uint64, error) {
	var resp campaign.CreateResponse
	err := clt.Post(ctx, "2/campaign/create/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.CampaignID, nil
}
