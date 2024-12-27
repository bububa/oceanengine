package sharedwallet

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/sharedwallet"
)

// WalletInfoGet 资金共享-批量查询钱包信息
// 批量查询钱包信息（包含共享钱包和子钱包）
func WalletInfoGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *sharedwallet.WalletInfoGetRequest) ([]sharedwallet.WalletInfo, error) {
	var resp sharedwallet.WalletInfoGetResponse
	if err := clt.GetAPI(ctx, "v3.0/shared_wallet/wallet_info/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.WalletInfo, nil
}
