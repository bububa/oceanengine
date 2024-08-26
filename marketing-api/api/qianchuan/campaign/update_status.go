package campaign

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/campaign"
)

// UpdateStatus 广告组更新状态
// 此接口用于更新当前千川广告账户下的广告组的状态。
// 本接口为批量更新接口，调用结果针对单个广告组存在部分成功部分失败场景，请避免根据应答code字段直接判断广告组状态更新的结果。
func UpdateStatus(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.UpdateStatusRequest) (*campaign.UpdateStatusResponseData, error) {
	var resp campaign.UpdateStatusResponse
	err := clt.Post(ctx, "v1.0/qianchuan/batch_campaign_status/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
