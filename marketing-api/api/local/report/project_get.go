package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/report"
)

// ProjectGet 获取项目数据
// 获取项目维度报表数据内容
func ProjectGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.ProjectGetRequest) (*report.ProjectGetResult, error) {
	var resp report.ProjectGetResponse
	if err := clt.GetAPI(ctx, "v3.0/local/report/project/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
