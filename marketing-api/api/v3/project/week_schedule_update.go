package project

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/project"
)

// WeekScheduleUpdate 批量更新项目投放时段
func WeekScheduleUpdate(clt *core.SDKClient, accessToken string, req *project.WeekScheduleUpdateRequest) (*project.UpdateResponseData, error) {
	var resp project.UpdateResponse
	if err := clt.Post("v3.0/project/week_schedule/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
