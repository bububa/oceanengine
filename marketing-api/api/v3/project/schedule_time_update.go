package project

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/project"
)

// ScheduleTimeUpdate 批量更新项目投放时间
func ScheduleTimeUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *project.ScheduleTimeUpdateRequest) (*project.UpdateResponseData, error) {
	var resp project.UpdateResponse
	if err := clt.PostAPI(ctx, "v3.0/project/schedule_time/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
