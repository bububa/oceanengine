package agent

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateStatementInvoiceRequest 开票-新建开票申请单（代理商版） API Request
type CreateStatementInvoiceRequest struct {
	// AgentIDs 代理商ID
	AgentIDs []uint64 `json:"agent_ids,omitempty"`
	// CustomerIDList 客户ID列表，鉴权使用
	CustomerIDList []uint64 `json:"customer_id_list,omitempty"`
	// StatementSerial 结算单编号
	StatementSerial string `json:"statement_serial,omitempty"`
	// RebateItemList 差额返点编号、金额集合
	RebateItemList []RebateItem `json:"rebate_item_list,omitempty"`
	// InvoiceType 开票类型 可选值:
	// GENERAL 增值税普通发票
	// SPECIAL 增值税专用发票
	// E_GENERAL 增值税电子普通发票
	// E_SPECIAL 增值税电子专用发票
	// ALL_E_GENERAL 电子发票（普通发票）
	// ALL_E_SPECIAL 电子发票（增值税专用发票）
	InvoiceType enum.InvoiceType `json:"invoice_type,omitempty"`
	// UnprintableRemark 非打印备注
	UnprintableRemark string `json:"unprintable_remark,omitempty"`
	// CustomerSubjectName 开票客户名称
	CustomerSubjectName string `json:"customer_subject_name,omitempty"`
	// CustomerTaxNo 纳税人识别号
	CustomerTaxNo string `json:"customer_tax_no,omitempty"`
	// CustomerAddress 企业注册地址
	CustomerAddress string `json:"customer_address,omitempty"`
	// CustomerPhone 企业注册电话
	CustomerPhone string `json:"customer_phone,omitempty"`
	// CustomerBank 企业开户银行
	CustomerBank string `json:"customer_bank,omitempty"`
	// CustomerBankAccount 企业银行账号
	CustomerBankAccount string `json:"customer_bank_account,omitempty"`
	// InvoiceBillList 发票票据列表
	InvoiceBillList []InvoiceBill `json:"invoice_bill_list,omitempty"`
	// SelectAddressAndPhone 数电发票，是否勾选企业注册地址、电话
	SelectAddressAndPhone bool `json:"select_address_and_phone,omitempty"`
	// SelectBankAndAccount 数电发票，是否勾选开户行、开户账号
	SelectBankAndAccount bool `json:"select_bank_and_account,omitempty"`
	// CustomerEmail 客户邮箱
	CustomerEmail string `json:"customer_email,omitempty"`
	// CustomerSmsPhone 客户短信电话
	CustomerSmsPhone string `json:"customer_sms_phone,omitempty"`
}

// RebateItem 差额返点编号、金额
type RebateItem struct {
	// RebateSerial 返点编号
	RebateSerial string `json:"rebate_serial,omitempty"`
	// ApplyAmount 返点申请金额
	ApplyAmount float64 `json:"apply_amount,omitempty"`
}

// InvoiceBill 发票票据
type InvoiceBill struct {
	// InvoiceBillProjectList 开票票据项目列表
	InvoiceBillProjectList []InvoiceBillProject `json:"invoice_bill_project_list,omitempty"`
	// Remark 打印备注
	Remark string `json:"remark,omitempty"`
}

// InvoiceBillProject 开票票据项目
type InvoiceBillProject struct {
	// InvoiceProjectName 开票项目名称
	InvoiceProjectName string `json:"invoice_project_name,omitempty"`
	// ApplyAmount 开票项目金额
	ApplyAmount float64 `json:"apply_amount,omitempty"`
	// Specifications 规则型号
	Specifications string `json:"specifications,omitempty"`
	// Unit 单位
	Unit string `json:"unit,omitempty"`
}

// Encode implements PostRequest interface
func (r CreateStatementInvoiceRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateStatementInvoiceResponse 开票-新建开票申请单（代理商版） API Response
type CreateStatementInvoiceResponse struct {
	model.BaseResponse
	Data struct {
		// InvoiceSerial 申请单编号
		InvoiceSerial string `json:"invoice_serial,omitempty"`
	} `json:"data,omitempty"`
}
