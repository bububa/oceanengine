package fund

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/fund"
)

// SharedWalletFundGet 获取共享钱包余额
// 返货相关需要咨询相关的运营和销售同学对接，具备返货相关前置条件下，相关返货资金信息可以通过本接口获得
func SharedWalletFundGet(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserIDs []uint64) ([]fund.SharedWalletBalance, error) {
	req := &fund.SharedWalletBalanceGetRequest{
		AdvertiserIDs: advertiserIDs,
	}
	var resp fund.SharedWalletBalanceGetResponse
	err := clt.Get(ctx, "2/fund/shared_wallet_balance/get", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
