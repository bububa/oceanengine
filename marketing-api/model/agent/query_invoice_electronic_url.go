package agent

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QueryInvoiceElectronicURLRequest 开票-获取电子发票文件接口（代理商版） API Request
type QueryInvoiceElectronicURLRequest struct {
	// AgentIDs 代理商ID
	AgentIDs []uint64 `json:"agent_ids,omitempty"`
	// InvoiceSerial 开票单编号
	InvoiceSerial string `json:"invoice_serial,omitempty"`
}

// Encode implements GetRequest interface
func (r QueryInvoiceElectronicURLRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("agent_ids", string(util.JSONMarshal(r.AgentIDs)))
	values.Set("invoice_serial", r.InvoiceSerial)
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QueryInvoiceElectronicURLResponse  开票-获取电子发票文件接口（代理商版）API Response
type QueryInvoiceElectronicURLResponse struct {
	model.BaseResponse
	Data struct {
		// URLList 电票下载链接
		URLList []InvoiceElectronicURL `json:"url_list,omitempty"`
	} `json:"data,omitempty"`
}

// InvoiceElectronicURL 电票下载链接
type InvoiceElectronicURL struct {
	// InvoiceNo 发票号码
	InvoiceNo string `json:"invoice_no,omitempty"`
	// InvoiceCode 发票代码
	InvoiceCode string `json:"invoice_code,omitempty"`
	// OfdURL 发票ofd链接
	OfdURL string `json:"ofd_url,omitempty"`
	// XmlURL 发票xml链接
	XmlURL string `json:"xml_url,omitempty"`
	// PdfURL 发票pdf链接
	PdfURL string `json:"pdf_url,omitempty"`
}
