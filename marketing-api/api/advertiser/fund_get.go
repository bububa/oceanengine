package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// FundGet 查询账号余额
// 获取广告主或代理商账户余额信息。
func FundGet(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) (*advertiser.FundGetResult, error) {
	req := &advertiser.FundGetRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.FundGetResponse
	err := clt.Get(ctx, "2/advertiser/fund/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
