package sharedwallet

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/sharedwallet"
)

// TransferCreate 资金共享-发起转账
// 发起转账，支持大钱包与小钱包互转，1:N批量转账
func TransferCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *sharedwallet.TransferCreateRequest) (string, error) {
	var resp sharedwallet.TransferCreateResponse
	if err := clt.PostAPI(ctx, "v3.0/cg_transfer/wallet/transfer/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.TransfrSerial, nil
}
