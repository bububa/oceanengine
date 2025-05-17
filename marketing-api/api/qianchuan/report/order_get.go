package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// OrderGet 获取随心推订单数据
// 此接口用户获取小店随心推投放订单数据报表，不支持获取PC广告投放数据
// 注意：如果您之前通过获取广告计划数据来拉取小店随心推订单数据报表，请尽快切换至该接口
func OrderGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponseData, error) {
	var resp report.GetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/aweme/report/order/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
