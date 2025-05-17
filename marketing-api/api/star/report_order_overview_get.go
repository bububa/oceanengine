package star

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/star"
)

// ReportOrderOverviewGet 获取订单投后分析报表
// 此接口用于根据订单id，获取星图客户所发布任务的接单达人作品的创意表现、传播表现、舆情表现、组件转化表现、订单性价比等效果分析类数据。
// 结果中所有的rate均为*100后的结果，如five_s_play_rate=2750，表示有效播放率27.5%。
// 数据为非实时更新，一般在次日凌晨产出前一天的数据；
// 一般历史数据都不会变，除了数据有问题有校对的情况会更新历史数据。
func ReportOrderOverviewGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *star.ReportOrderOverviewGetRequest) (*star.ReportOrderOverviewGetResult, error) {
	var resp star.ReportOrderOverviewGetResponse
	if err := clt.GetAPI(ctx, "2/star/report/order_overview/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
