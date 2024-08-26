package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// RoiGoalUpdate 更新计划的支付ROI目标
// 此接口用于批量更新广告计划的支付ROI目标
func RoiGoalUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.RoiGoalUpdateRequest) ([]ad.RoiGoalUpdateResult, error) {
	var resp ad.RoiGoalUpdateResponse
	if err := clt.Post(ctx, "v1.0/qianchuan/roi/goal/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Results, nil
}
