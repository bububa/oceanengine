package unipromotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	unipromotion "github.com/bububa/oceanengine/marketing-api/model/qianchuan/uni_promotion"
)

// OrderReportGet 获取随心推全域订单数据
func OrderReportGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *unipromotion.OrderReportGetRequest) (*unipromotion.OrderReportGetResult, error) {
	var resp unipromotion.OrderReportGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/aweme/uni_promotion/order/report/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
