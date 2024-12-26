package customercenter

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/customercenter"
)

// TransferBalanceGet 工作台转账-查询账户转账余额
// 查询账户自身转账余额、作为转出方需要预留的竞价消耗保证金
func TransferBalanceGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *customercenter.TransferBalanceGetRequest) ([]customercenter.TargetAmountDetail, error) {
	var resp customercenter.TransferBalanceGetResponse
	if err := clt.GetAPI(ctx, "v3.0/cg_transfer/transfer_balance/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.TargetAmountDetailList, nil
}
