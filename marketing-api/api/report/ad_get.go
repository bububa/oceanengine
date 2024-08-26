package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report"
)

// AdGet 广告计划数据
// 此接口用于获取广告计划纬度的投放数据，包括消耗、点击、展示等指标，具体可以参考应答参数指标说明。
// 默认不获取删除广告计划数据，如果需要查询删除的数据，可以在filtering中设置status，传AD_STATUS_ALL（所有包含已删除）
func AdGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponseData, error) {
	var resp report.GetResponse
	err := clt.Get(ctx, "2/report/ad/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
