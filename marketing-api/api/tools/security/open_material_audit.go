package security

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/security"
)

// OpenMaterialAudit 广告素材预审
// 与广告素材使用相同规则进行素材预审，检验素材是否符合广告投放审核要求
func OpenMaterialAudit(ctx context.Context, clt *core.SDKClient, accessToken string, req *security.OpenMaterialAuditRequest) (*security.OpenMaterialAuditResult, error) {
	var resp security.OpenMaterialAuditResponse
	if err := clt.PostAPI(ctx, "v3.0/security/open_material_audit/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
