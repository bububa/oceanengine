package live

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/live"
)

// RoomDetailGet 获取直播间详情
func RoomDetailGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *live.RoomDetailGetRequest) (*live.Room, error) {
	var resp live.RoomDetailGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/today_live/room/detail/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
