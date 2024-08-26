package enterprise

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/enterprise"
)

// FlowCategoryGet 获取企业号流量来源数据
func FlowCategoryGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *enterprise.ReportRequest) (*enterprise.ReportResult, error) {
	var resp enterprise.ReportResponse
	if err := clt.Get(ctx, "v1.0/enterprise/flow/category/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
