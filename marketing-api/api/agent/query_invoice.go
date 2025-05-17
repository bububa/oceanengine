package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// QueryInvoice 开票-查询开票单数据（代理商版）
// 通过结算单/项目查询开票单数据
func QueryInvoice(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.QueryInvoiceRequest) (*agent.QueryInvoiceResult, error) {
	var resp agent.QueryInvoiceResponse
	if err := clt.GetAPI(ctx, "2/query/invoice/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
