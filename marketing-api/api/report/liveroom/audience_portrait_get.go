package liveroom

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report/liveroom"
)

// AudiencePortraitGet 直播受众分析报表
func AudiencePortraitGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *liveroom.AudiencePortraitGetRequest) (*liveroom.AudiencePortraitGetResult, error) {
	var resp liveroom.AudiencePortraitGetResponse
	if err := clt.Get(ctx, "v3.0/report/report/live_room/audience/portrait/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
