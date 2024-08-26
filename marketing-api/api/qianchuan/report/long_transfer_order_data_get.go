package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// LongTransferOrderDataGet 获取长周期订单数据
func LongTransferOrderDataGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.CustomGetRequest) (*report.CustomGetResult, error) {
	var resp report.CustomGetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/report/long_transfer/order/data/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
