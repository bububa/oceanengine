package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// MaterialGet 获取广告素材数据
func MaterialGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponseData, error) {
	var resp report.GetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/report/material/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
