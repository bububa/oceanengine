package agent

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QueryInvoiceRequest 开票-查询开票单数据（代理商版） API Request
type QueryInvoiceRequest struct {
	// AgentID 代理商ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// Filtering 过滤器
	Filtering *QueryInvoiceFilter `json:"filtering,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 每页数量，最多 100
	PageSize int `json:"page_size,omitempty"`
}

type QueryInvoiceFilter struct {
	// StatementSerials 结算单编号
	StatementSerials []string `json:"statement_serials,omitempty"`
	// ProjectSerials 项目编号
	ProjectSerials []string `json:"project_serials,omitempty"`
	// InvoiceStatuses 开票状态 0：已作废，1：未开票，2：部分开票，3：已开票，4：无需开票
	InvoiceStatuses []int `json:"invoice_statuses,omitempty"`
	// InvoiceSerialList 开票单编号
	InvoiceSerialList []string `json:"invoice_serial_list,omitempty"`
	// ContractSerial 合同编号
	ContractSerial string `json:"contract_serial,omitempty"`
	// SubmitStartTime 提交审批时间范围开始时间，格式：%Y-%m-%d %H:%M:%S
	SubmitStartTime string `json:"submit_start_time,omitempty"`
	// SubmitEndTime 提交审批时间范围结束时间，格式：%Y-%m-%d %H:%M:%S
	SubmitEndTime string `json:"submit_end_time,omitempty"`
	// InvoiceStartDate 实际开票时间范围开始时间，格式：%Y-%m-%d
	InvoiceStartDate string `json:"invoice_start_date,omitempty"`
	// InvoiceEndDate 实际开票时间范围结束时间，格式：%Y-%m-%d
	InvoiceEndDate string `json:"invoice_end_date,omitempty"`
	// InvoiceType 发票类型 可选值:
	// GENERAL 增值税普通发票
	// SPECIAL 增值税专用发票
	// E_GENERAL 增值税电子普通发票
	// E_SPECIAL 增值税电子专用发票
	// ALL_E_GENERAL 电子发票（普通发票）
	// ALL_E_SPECIAL 电子发票（增值税专用发票）
	InvoiceType enum.InvoiceType `json:"invoice_type,omitempty"`
	// DifferenceInvoice 是否差额开票 可选值:
	// FULL 全额开票
	// DIFF 差额开票
	DifferenceInvoice enum.DifferenceInvoice `json:"difference_invoice,omitempty"`
	// RevertStatusList 红冲状态 可选值:
	// 1 未红冲
	// 2 全额红冲
	// 3 红票
	// 4 部分红冲
	RevertStatusList []int `json:"revert_status_list,omitempty"`
	// Platform 开票平台 可选值:
	// AD 广告
	// QC 千川
	// BDT 本地推
	Platform string `json:"platform,omitempty"`
}

// Encode implements GetRequest interface
func (r QueryInvoiceRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("agent_id", strconv.FormatUint(r.AgentID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QueryInvoiceResponse 开票-查询开票单数据（代理商版） API Response
type QueryInvoiceResponse struct {
	model.BaseResponse
	Data *QueryInvoiceResult `json:"data,omitempty"`
}

type QueryInvoiceResult struct {
	// PageInfo 翻页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// InvoiceInfoList   v
	InvoiceInfoList []InvoiceInfo `json:"invoice_info_list,omitempty"`
}

type InvoiceInfo struct {
	// InvoiceID 开票单ID
	InvoiceID uint64 `json:"invoice_id,omitempty"`
	// InvoiceSerial 开票单编号
	InvoiceSerial string `json:"invoice_serial,omitempty"`
	// InvoiceStatus 开票状态0：已作废，1：未开票，2：部分开票，3：已开票，4：无需开票
	InvoiceStatus int `json:"invoice_status,omitempty"`
	// InvoiceStatusName 开票状态名称
	InvoiceStatusName string `json:"invoice_status_name,omitempty"`
	// InvoiceAmount 开票金额
	InvoiceAmount float64 `json:"invoice_amount,omitempty"`
	// AbandonedAmount 作废金额
	AbandonedAmount float64 `json:"abandoned_amount,omitempty"`
	// ActualInvoiceAmount 实际开票金额
	ActualInvoiceAmount float64 `json:"actual_invoice_amount,omitempty"`
	// ApplyAmount 申请金额
	ApplyAmount float64 `json:"apply_amount,omitempty"`
	// IssueTime 发票开具时间
	IssueTime string `json:"issue_time,omitempty"`
	// RevertStatus 红冲状态 1：未红冲，2：全额红冲，3：红票，4：部分红冲
	RevertStatus int `json:"revert_status,omitempty"`
	// RevertStatusName 红冲状态名称
	RevertStatusName string `json:"revert_status_name,omitempty"`
	// InvoiceNos 发票号码
	InvoiceNos []string `json:"invoice_nos,omitempty"`
	// InvoiceCodes 发票代码
	InvoiceCodes []string `json:"invoice_codes,omitempty"`
	// RevertFrozenAmount 红冲冻结金额
	RevertFrozenAmount float64 `json:"revert_frozen_amount,omitempty"`
	// RevertAmount 红冲金额
	RevertAmount float64 `json:"revert_amount,omitempty"`
	// InvoiceType 发票类型ID
	InvoiceType int `json:"invoice_type,omitempty"`
	// InvoiceTypeName 发票类型名称
	InvoiceTypeName string `json:"invoice_type_name,omitempty"`
	// ContractSerial 合同编号
	ContractSerial string `json:"contract_serialon,omitempty"`
	// ContractName 合同名称                              t
	ContractName string `json:"contract_name,omitempty"`
	// ContractSubjectName 媒体主体名称
	ContractSubjectName string `json:"contract_subject_name,omitempty"`
	// CustomerID 客户ID
	CustomerID uint64 `json:"customer_id,omitempty"`
	// CustomerName 代理主体名称
	CustomerName string `json:"customer_name,omitempty"`
	// PlatformName 业务线名称
	PlatformName string `json:"platform_name,omitempty"`
	// InvoiceObjectName 开票类型名称
	InvoiceObjectName string `json:"invoice_object_name,omitempty"`
	// CustomerSubjectName 开票客户名称
	CustomerSubjectName string `json:"customer_subject_name,omitempty"`
	// InvoiceProjectList 开票项目名称
	InvoiceProjectList []string `json:"invoice_project_list,omitempty"`
	// SubmitterName 提交人名称
	SubmitterName string `json:"submitter_name,omitempty"`
	// SubmitTime 提交时间 yyyy-MM-dd HH:mm:ss          u
	SubmitTime string `json:"submit_time,omitempty"`
	// DifferenceInvoiceName 是否差额开票
	DifferenceInvoiceName string `json:"difference_invoice_name,omitempty"`
}
