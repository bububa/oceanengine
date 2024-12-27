package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// QueryRebateAccountingInfo 返点-查询返点核算流水
// 结算-查询返点核算信息
func QueryRebateAccountingInfo(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.QueryRebateAccountingInfoRequest) (*agent.QueryRebateAccountingInfoResult, error) {
	var resp agent.QueryRebateAccountingInfoResponse
	if err := clt.GetAPI(ctx, "2/query/rebate_accounting_info/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
