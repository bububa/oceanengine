package advertiser

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// AccountFundGet 批量查询账户余额
// 批量查询账号余额，支持AD、星图、本地推账户
func AccountFundGet(clt *core.SDKClient, accessToken string, req *advertiser.AccountFundGetRequest) ([]advertiser.AccountFund, error) {
	var resp advertiser.AccountFundGetResponse
	if err := clt.Get("v3.0/account/fund/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
