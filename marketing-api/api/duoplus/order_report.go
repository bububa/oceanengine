package duoplus

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/duoplus"
)

// OrderReport 获取订单数据报表
// 获取DOU+订单指标数据
func OrderReport(clt *core.SDKClient, accessToken string, req *duoplus.OrderReportRequest) (*duoplus.OrderReportResult, error) {
	var resp duoplus.OrderReportResponse
	if err := clt.Get("v3.0/duoplus/order/report/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
