package live

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/live"
)

// RoomGet 获取今日直播间列表
// 支持时间筛选范围内直播间列表
func RoomGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *live.RoomGetRequest) (*live.RoomGetData, error) {
	var resp live.RoomGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/today_live/room/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
