package sharedwallet

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/sharedwallet"
)

// TransactionDetailGet  资金共享-查询共享钱包流水明细
func TransactionDetailGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *sharedwallet.TransactionDetailGetRequest) (*sharedwallet.TransactionDetailGetResult, error) {
	var resp sharedwallet.TransactionDetailGetResponse
	if err := clt.GetAPI(ctx, "v3.0/shared_wallet/transaction_detail/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
