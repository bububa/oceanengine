package sharedwallet

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/sharedwallet"
)

// WalletBalanceGet  资金共享-批量查询钱包余额
func WalletBalanceGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *sharedwallet.WalletBalanceGetRequest) (*sharedwallet.WalletBalanceInfo, error) {
	var resp sharedwallet.WalletBalanceGetResponse
	if err := clt.GetAPI(ctx, "v3.0/shared_wallet/wallet_balance/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
