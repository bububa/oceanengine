package live

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/live"
)

// RoomUserGet 获取直播间用户洞察
func RoomUserGet(clt *core.SDKClient, accessToken string, req *live.RoomUserGetRequest) (*live.RoomUser, error) {
	var resp live.RoomUserGetResponse
	if err := clt.Get("v1.0/qianchuan/today_live/room/user/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
