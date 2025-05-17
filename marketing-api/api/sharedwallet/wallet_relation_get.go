package sharedwallet

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/sharedwallet"
)

// WalletRelationGet 资金共享-查询子钱包下绑定的adv列表
// 查询子钱包下绑定的adv列表, 支持分页
func WalletRelationGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *sharedwallet.WalletRelationGetRequest) (*sharedwallet.WalletRelationGetResult, error) {
	var resp sharedwallet.WalletRelationGetResponse
	if err := clt.GetAPI(ctx, "v3.0/shared_wallet/wallet_relation/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
