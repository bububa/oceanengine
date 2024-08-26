package liveroom

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report/liveroom"
)

// FlowCategoryGet 直播间流量来源报表
// 此接口用于获取直播间的流量来源数据；
// 流量来源包含竞价广告、品牌广告、DOU+流量、自然流量；
// 支持获取分流量来源的直播间观看数、观看人数、人均停留时长、关注数、打赏次数、礼物总金额、商品订单数、商品订单金额；
func FlowCategoryGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *liveroom.Request) (*liveroom.ResponseData, error) {
	var resp liveroom.Response
	err := clt.Get(ctx, "2/report/live_room/flow_category/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
