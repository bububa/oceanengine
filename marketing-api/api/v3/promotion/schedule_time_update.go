package promotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/promotion"
)

// ScheduleTimeUpdate 批量更新广告投放时段
// 支持在广告层级下批量修改投放时段，所设置的广告投放时段必须在广告所属项目的投放时段范围内
// 仅支持手动投放的广告
func ScheduleTimeUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *promotion.ScheduleTimeUpdateRequest) (*promotion.UpdateResponseData, error) {
	var resp promotion.UpdateResponse
	if err := clt.PostAPI(ctx, "v3.0/promotion/schedule_time/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
