package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// ScheduleTimeUpdate 更新计划投放时段
// 批量更新计划投放时段，对于live_schedule_type=SCHEDULE_TIME_FIXEDRANGE的计划不支持
func ScheduleTimeUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.ScheduleTimeUpdateRequest) (*ad.UpdateResponseData, error) {
	var resp ad.UpdateResponse
	if err := clt.Post(ctx, "v1.0/qianchuan/ad/schedule_time/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
