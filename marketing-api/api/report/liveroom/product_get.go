package liveroom

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report/liveroom"
)

// ProductGet 直播间商品分析报表
// 此接口用于获取直播间的商品分析数据，包含在售商品信息和商品转化数据；
// 在售商品信息：商品id、商品名称、商品图片url、商品落地页url；
// 商品转化数据：商品点击数、商品点击人数、商品下单数、商品下单人数、商品订单数、商品订单人数、商品订单金额；
func ProductGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *liveroom.Request) (*liveroom.ResponseData, error) {
	var resp liveroom.Response
	err := clt.Get(ctx, "2/report/live_room/product/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
