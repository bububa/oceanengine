package businessplatform

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/businessplatform"
)

// CompanyInfoGet 获取纵横组织下所有主体信息
func CompanyInfoGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *businessplatform.CompanyInfoGetRequest) (*businessplatform.CompanyInfoGetData, error) {
	var resp businessplatform.CompanyInfoGetResponse
	if err := clt.GetAPI(ctx, "v3.0/business_platform/company_info/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
