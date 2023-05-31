package report

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// LongTransferOrderGet 长周期转化价值-订单明细
func LongTransferOrderGet(clt *core.SDKClient, accessToken string, req *report.LongTransferOrderGetRequest) (*report.LongTransferOrderGetData, error) {
	var resp report.LongTransferOrderGetResponse
	if err := clt.Get("v1.0/qianchuan/report/long_transfer/order/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
