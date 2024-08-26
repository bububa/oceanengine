package liveroom

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report/liveroom"
)

// AttributeGet 直播间属性报表
// 此接口用于通过广告主id获取广告主账号绑定的抖音号的开播信息，包含主播信息和直播间信息；
// 主播信息：主播id、主播头像url、主播昵称；
// 直播间信息：直播间id、直播间标题、直播间封面url、直播开始时间、直播结束时间、直播间状态、直播间二维码url；
func AttributeGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *liveroom.Request) (*liveroom.ResponseData, error) {
	var resp liveroom.Response
	err := clt.Get(ctx, "2/report/live_room/attribute/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
