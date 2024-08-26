package campaign

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/campaign"
)

// Update 修改广告组
// 此接口用于修改千川广告账户下的广告组
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.UpdateRequest) (uint64, error) {
	var resp campaign.CreateResponse
	err := clt.Post(ctx, "v1.0/qianchuan/campaign/update/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.CampaignID, nil
}
