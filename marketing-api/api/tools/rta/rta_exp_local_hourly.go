package rta

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/rta"
)

// RtaExpLocalHourlyGet 获取站内媒体RTA联合实验数据（分时t+5）
// 功能
// 该接口用于查询站内媒体渠道的RTA联合实验数据，支持分时t+5级别数据
func RtaExpLocalHourlyGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *rta.RtaExpLocalHourlyGetRequest) ([]rta.Report, error) {
	var resp rta.RtaExpLocalHourlyGetResponse
	if err := clt.GetAPI(ctx, "v3.0/report/rta_exp_local_hourly/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Data, nil
}
