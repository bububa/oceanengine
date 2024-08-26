package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// CreativeGet 广告创意数据
// 目前批量拉取多个创意数据返回时，若所有字段均无值，则该条创意不返回，即有可能拉取10条创意，仅返回7条创意的数据，未返回数据的创意，可当做无数据
// 目前暂不支持拉取开启“计划托管”功能的广告创意数据
func CreativeGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponseData, error) {
	var resp report.GetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/report/creative/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
