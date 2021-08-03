package liveroom

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report/liveroom"
)

// AnalysisGet 直播间分析报表
// 互动分析数据：评论数、分享数、关注数、打赏次数、礼物总金额等；
// 商品转化数据：查看购物车数、商品点击数、商品点击人数、商品下单数、商品下单人数、商品订单数、商品订单人数、商品订单金额等；
func AnalysisGet(clt *core.SDKClient, accessToken string, req *liveroom.Request) (*liveroom.ResponseData, error) {
	var resp liveroom.Response
	err := clt.Get("2/report/live_room/analysis/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
