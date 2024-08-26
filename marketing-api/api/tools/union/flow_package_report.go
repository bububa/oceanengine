package union

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/union"
)

// FlowPackageReport 查看rit数据
// 查看穿山甲广告位数据，包括付费率、消耗等数据，默认查询历史30天数据。
func FlowPackageReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *union.FlowPackageReportRequest) (*union.FlowPackageReportData, error) {
	var resp union.FlowPackageReportResponse
	if err := clt.Get(ctx, "2/tools/union/flow_package/report/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
