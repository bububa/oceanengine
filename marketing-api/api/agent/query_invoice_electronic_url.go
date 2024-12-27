package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// QueryInvoiceElectronicURL 开票-获取电子发票文件接口（代理商版）
// 开票-获取电票/数电票发票文件下载链接接口
func QueryInvoiceElectronicURL(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.QueryInvoiceElectronicURLRequest) ([]agent.InvoiceElectronicURL, error) {
	var resp agent.QueryInvoiceElectronicURLResponse
	if err := clt.GetAPI(ctx, "2/query/invoice_electronic_url/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.URLList, nil
}
