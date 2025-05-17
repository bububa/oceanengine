package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// QueryTransferBalance 转账-查询账户转账余额（代理）
// 查询账户自身转账余额、作为减款方需要预留的竞价消耗保证金
func QueryTransferBalance(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.QueryTransferBalanceRequest) ([]agent.AccountAmountDetail, error) {
	var resp agent.QueryTransferBalanceResponse
	if err := clt.GetAPI(ctx, "v3.0/cg_transfer/query_transfer_balance/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.AccountAmountDetailList, nil
}
