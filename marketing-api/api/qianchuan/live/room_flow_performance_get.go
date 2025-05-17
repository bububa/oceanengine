package live

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/live"
)

// RoomFlowPerformanceGet 获取直播间流量表现
// 支持获取直播间流量表现，对应PC今日直播大屏数据
func RoomFlowPerformanceGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *live.RoomFlowPerformanceGetRequest) (*live.RoomFlowPerformance, error) {
	var resp live.RoomFlowPerformanceGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/today_live/room/flow_performance/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
