package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// TodayLiveRoomConfigGet 获取直播大屏可用指标和维度
func TodayLiveRoomConfigGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.CustomConfigGetRequest) (*report.CustomConfigGetResult, error) {
	var resp report.CustomConfigGetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/report/today_live/room/config/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
