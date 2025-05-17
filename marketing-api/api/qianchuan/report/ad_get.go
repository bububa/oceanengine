package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// AdGet 广告计划数据
// 目前批量拉取多个计划数据返回时，若所有字段均无值，则该条计划不返回，即有可能拉取10条计划，仅返回7条计划的数据，未返回数据的计划，可当做无数据
// 目前暂不支持拉取开启“计划托管”功能的广告计划数据
func AdGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponseData, error) {
	var resp report.GetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/report/ad/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
