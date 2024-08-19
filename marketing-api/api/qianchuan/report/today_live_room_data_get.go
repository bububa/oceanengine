package report

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// TodayLiveRoomDataGet 获取直播大屏数据
func TodayLiveRoomDataGet(clt *core.SDKClient, accessToken string, req *report.CustomGetRequest) (*report.CustomGetResult, error) {
	var resp report.CustomGetResponse
	err := clt.Get("v1.0/qianchuan/report/today_live/room/data/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
