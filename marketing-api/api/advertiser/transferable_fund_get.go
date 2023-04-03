package advertiser

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// TransferableFundGet 查询账户可转余额
// 通过此接口可获取广告主、代理商或组织账户下的可转资金余额信息。
func TransferableFundGet(clt *core.SDKClient, accessToken string, advertiserID uint64) (*advertiser.TransferableFund, error) {
	req := &advertiser.TransferableFundGetRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.TransferableFundGetResponse
	err := clt.Get("2/advertiser/transferable_fund/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
