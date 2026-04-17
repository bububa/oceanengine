package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/report"
)

// AccountGet 查询账户数据
// 查询本地推账户维度数据
func AccountGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.AccountGetRequest) (*report.AccountGetResult, error) {
	var resp report.AccountGetResponse
	if err := clt.GetAPI(ctx, "v3.0/local/report/account/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
