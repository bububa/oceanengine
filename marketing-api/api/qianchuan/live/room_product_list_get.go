package live

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/live"
)

// RoomProductListGet 获取直播间商品列表
func RoomProductListGet(clt *core.SDKClient, accessToken string, req *live.RoomProductListGetRequest) (*live.RoomProductListData, error) {
	var resp live.RoomProductListGetResponse
	if err := clt.Get("v1.0/qianchuan/today_live/room/product_list/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
