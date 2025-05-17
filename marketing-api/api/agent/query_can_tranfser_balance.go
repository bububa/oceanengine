package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// QueryCanTransferBalance 转账-获取最大可转余额（代理）
// 查询减款方与加款方之间最大可转金额，接口内已自动扣除需要预留的竞价消耗保证金，支持查询1:N转账的最大可转金额
func QueryCanTransferBalance(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.QueryCanTransferBalanceRequest) ([]agent.CanTransferDetail, error) {
	var resp agent.QueryCanTransferBalanceResponse
	if err := clt.GetAPI(ctx, "v3.0/cg_transfer/query_can_transfer_balance/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.CanTransferDetailList, nil
}
