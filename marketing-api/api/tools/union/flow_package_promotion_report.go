package union

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/union"
)

// FlowPackagePromotionReport 查看2.0rit数据
// 查看穿山甲2.0广告位数据
func FlowPackagePromotionReport(clt *core.SDKClient, accessToken string, req *union.FlowPackagePromotionReportRequest) (*union.FlowPackageReportData, error) {
	var resp union.FlowPackageReportResponse
	if err := clt.GetAPI("v3.0/tools/union/flow_package/promotion/report/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
