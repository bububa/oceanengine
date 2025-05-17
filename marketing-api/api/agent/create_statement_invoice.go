package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// CreateStatementInvoice 开票-新建开票申请单（代理商版）
// 创建结算单开票接口，支持差额开票
func CreateStatementInvoice(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.CreateStatementInvoiceRequest) (string, error) {
	var resp agent.CreateStatementInvoiceResponse
	if err := clt.PostAPI(ctx, "2/create/statement_invoice/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.InvoiceSerial, nil
}
