package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/report"
)

// MaterialGet 获取素材数据
// 获取素材维度报表数据内容获取
func MaterialGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.MaterialGetRequest) (*report.MaterialGetResult, error) {
	var resp report.MaterialGetResponse
	if err := clt.GetAPI(ctx, "v3.0/local/report/material/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
