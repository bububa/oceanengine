package sharedwallet

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/sharedwallet"
)

// MainWalletGet 资金共享-共享钱包信息查询
// 查询当前共享钱包(大钱包)的信息
func MainWalletGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *sharedwallet.MainWalletGetRequest) (*sharedwallet.MainWalletInfo, error) {
	var resp sharedwallet.MainWalletGetResponse
	if err := clt.GetAPI(ctx, "v3.0/shared_wallet/main_wallet/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
