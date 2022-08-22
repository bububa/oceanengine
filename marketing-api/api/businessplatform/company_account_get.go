package businessplatform

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/businessplatform"
)

// CompanyAccountGet 获取主体下的账户列表
func CompanyAccountGet(clt *core.SDKClient, accessToken string, req *businessplatform.CompanyAccountGetRequest) (*businessplatform.CompanyAccountGetData, error) {
	var resp businessplatform.CompanyAccountGetResponse
	if err := clt.Get("v3.0/business_platform/company_account/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
