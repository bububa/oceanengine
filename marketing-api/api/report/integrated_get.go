package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report"
)

// IntegratedGet 多合一数据报表接口
// 数据报表使用方式：报表、监控通知、自动化调整
func IntegratedGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.IntegratedRequest) (*report.IntegratedResponseData, error) {
	var resp report.IntegratedResponse
	err := clt.Get(ctx, "2/report/integrated/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
