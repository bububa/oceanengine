package sharedwallet

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/sharedwallet"
)

// TransferList 资金共享-查询转账列表
// 通过筛选条件捞取转账记录
func TransferList(ctx context.Context, clt *core.SDKClient, accessToken string, req *sharedwallet.TransferListRequest) ([]sharedwallet.TransferRecord, error) {
	var resp sharedwallet.TransferListResponse
	if err := clt.GetAPI(ctx, "v3.0/cg_transfer/wallet/transfer/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.RecordList, nil
}
