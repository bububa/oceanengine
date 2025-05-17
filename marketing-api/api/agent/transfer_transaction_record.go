package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// TransferTransactionRecord 查询代理商转账记录
// 代理商转账记录查询，相关功能与代理商平台的「商务-转账记录-账户转账记录」模块对齐。
func TransferTransactionRecord(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.TransferTransactionRecordRequest) (*agent.TransferTransactionRecordResult, error) {
	var resp agent.TransferTransactionRecordResponse
	if err := clt.Get(ctx, "2/agent/transfer/transaction_record/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
