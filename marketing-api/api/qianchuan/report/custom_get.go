package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// CustomGet 自定义报表
// 自定义报表是基于千川自定义报表对于数据报表提供新的接口能力，用户可以自由选择和组合指标和维度，以此定义数据报表的字段。
// 相关指标和维度可通过获取自定义报表可用维度和指标接口获取不同数据主题下的可用维度和指标。
func CustomGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.CustomGetRequest) (*report.CustomGetResult, error) {
	var resp report.CustomGetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/report/custom/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
