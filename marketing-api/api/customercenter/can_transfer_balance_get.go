package customercenter

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/customercenter"
)

// CanTransferBalanceGet 工作台转账-获取最大可转余额
// 查询转出方与转入方之间最大可转金额，接口内已自动扣除需要预留的竞价消耗保证金，支持查询1:N转账的最大可转金额
func CanTransferBalanceGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *customercenter.CanTransferBalanceGetRequest) ([]customercenter.CanTransferDetail, error) {
	var resp customercenter.CanTransferBalanceGetResponse
	if err := clt.GetAPI(ctx, "v3.0/cg_transfer/can_transfer_balance/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.CanTransferDetailList, nil
}
