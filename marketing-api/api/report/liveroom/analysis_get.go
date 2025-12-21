package liveroom

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report/liveroom"
)

// AnalysisGet 直播间分析报表
// 直播间分析报表接口用于获取直播间的基础分析数据、互动分析数据
// 数据范围&口径：
// 仅支持获取广告主账户关联的抖音号的直播数据，不限制广告主账户和抖音号两者需为同一公司主体
// 仅支持查询2020年7月1日之后的数据
// 指标为非实时数据，一般在次日凌晨产出前一天的数据（与直播是否结束无关）
func AnalysisGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *liveroom.AnalysisGetRequest) (*liveroom.AnalysisGetResult, error) {
	var resp liveroom.AnalysisGetResponse
	if err := clt.Get(ctx, "v3.0/report/live_room/analysis/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
