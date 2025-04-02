package security

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/security"
)

// AuditResults 广告素材预审结果
// 拉取送审对象ID的审核结果
func AuditResults(ctx context.Context, clt *core.SDKClient, accessToken string, req *security.AuditResultsRequest) (*security.AuditResult, error) {
	var resp security.AuditResultsResponse
	if err := clt.GetAPI(ctx, "v3.0/security/audit_results/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
