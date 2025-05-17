package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// QueryRebateBalance 返点-查询返点流水
// 结算-查询返点流水信息
func QueryRebateBalance(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.QueryRebateBalanceRequest) (*agent.QueryRebateBalanceResult, error) {
	var resp agent.QueryRebateBalanceResponse
	if err := clt.GetAPI(ctx, "2/query/rebate_balance/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
