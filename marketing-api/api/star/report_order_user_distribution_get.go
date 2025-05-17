package star

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/star"
)

// ReportOrderUserDistributionGet 获取订单投后受众报表
// 此接口用于根据订单id，获取星图客户所发布任务的接单达人作品的受众分析数据，包含性别、年龄、省份、城市、设备、兴趣、活跃度分布。
// 数据为非实时更新，一般在次日凌晨产出前一天的数据；
// 一般历史数据都不会变，除了数据有问题有校对的情况会更新历史数据。
func ReportOrderUserDistributionGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *star.ReportOrderUserDistributionGetRequest) (*star.ReportOrderUserDistributionGetResult, error) {
	var resp star.ReportOrderUserDistributionGetResponse
	if err := clt.GetAPI(ctx, "2/star/report/order_user_distribution/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
