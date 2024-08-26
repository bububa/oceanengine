package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// QueryTransferDetail 转账-查询转账单信息（代理）
// 转账单信息，包括状态、双方账户、转账金额
func QueryTransferDetail(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.QueryTransferDetailRequest) (*agent.TransferDetail, error) {
	var resp agent.QueryTransferDetailResponse
	if err := clt.GetAPI(ctx, "v3.0/cg_transfer/query_transfer_detail/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
