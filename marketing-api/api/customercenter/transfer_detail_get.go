package customercenter

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/customercenter"
)

// TransferDetailGet 工作台转账-查询转账单信息
// 转账单信息，包括状态、双方账户、转账金额
func TransferDetailGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *customercenter.TransferDetailGetRequest) (*customercenter.TransferDetail, error) {
	var resp customercenter.TransferDetailGetResponse
	if err := clt.GetAPI(ctx, "v3.0/cg_transfer/transfer_detail/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
