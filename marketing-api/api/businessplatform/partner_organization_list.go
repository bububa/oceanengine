package businessplatform

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/businessplatform"
)

// PartnerOrganizationList 查询合作组织
func PartnerOrganizationList(ctx context.Context, clt *core.SDKClient, accessToken string, req *businessplatform.PartnerOrganizationListRequest) (*businessplatform.PartnerOrganizationListData, error) {
	var resp businessplatform.PartnerOrganizationListResponse
	if err := clt.Get(ctx, "2/business_platform/partner_organization/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
