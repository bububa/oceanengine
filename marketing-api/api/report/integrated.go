package report

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report"
)

// 多合一数据报表接口
// 数据报表使用方式：报表、监控通知、自动化调整
func Integrated(clt *core.SDKClient, accessToken string, req *report.IntegratedRequest) (*report.IntegratedResponse, error) {
	var resp report.IntegratedResponse
	err := clt.Get("2/report/integrated/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
