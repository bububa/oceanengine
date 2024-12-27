package sharedwallet

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/sharedwallet"
)

// TransferDetail 资金共享-查询转账单信息
// 转账单信息，包括状态、转账钱包id、转账金额
func TransferDetail(ctx context.Context, clt *core.SDKClient, accessToken string, req *sharedwallet.TransferDetailRequest) (*sharedwallet.TransferDetail, error) {
	var resp sharedwallet.TransferDetailResponse
	if err := clt.GetAPI(ctx, "v3.0/cg_transfer/wallet/transfer/detail/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
