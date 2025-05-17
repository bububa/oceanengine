package rta

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/rta"
)

// RtaExpLocalDailyGet 获取站内媒体RTA联合实验数据分天(t+1）
// 功能
// 该接口用于查询站内媒体渠道的RTA联合实验数据，支持分天t+1级别数据
// 注：由于数据更新时间存在波动性，建议在查询当日上午7点后尝试拉取前一天的数据
func RtaExpLocalDailyGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *rta.RtaExpLocalDailyGetRequest) ([]rta.Report, error) {
	var resp rta.RtaExpLocalDailyGetResponse
	if err := clt.GetAPI(ctx, "v3.0/report/rta_exp_local_daily/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Data, nil
}
