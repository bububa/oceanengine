package project

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/project"
)

// RoiGoalUpdate 批量修改项目ROI系数
// 本接口支持批量修改项目ROI系数，当前仅以下项目支持ROI系数修改：
// 应用推广、自动投放项目
// 电商推广、自动投放项目
// 电商推广、自动投放、周期稳投项目
func RoiGoalUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *project.RoiGoalUpdateRequest) (*project.UpdateResponseData, error) {
	var resp project.UpdateResponse
	if err := clt.PostAPI(ctx, "v3.0/project/roigoal/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
