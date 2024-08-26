package enterprise

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/enterprise"
)

// VideoInfoGet 获取企业号视频分析数据
func VideoInfoGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *enterprise.ReportRequest) (*enterprise.ReportResult, error) {
	var resp enterprise.ReportResponse
	if err := clt.Get(ctx, "v1.0/enterprise/video/info/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
