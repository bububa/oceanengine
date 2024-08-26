package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// ScheduleFixedRangeUpdate 更新计划投放时长
// 批量更新计划固定时长，仅支持marketing_goal=LIVE_PROM_GOODS&live_schedule_type=SCHEDULE_TIME_FIXEDRANGE计划支持
func ScheduleFixedRangeUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.ScheduleFixedRangeUpdateRequest) (*ad.UpdateResponseData, error) {
	var resp ad.UpdateResponse
	if err := clt.Post(ctx, "v1.0/qianchuan/ad/schedule_fixed_range/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
