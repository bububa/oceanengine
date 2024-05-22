package agent

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TransferTransactionRecordRequest 查询代理商转账记录 API Request
type TransferTransactionRecordRequest struct {
	// AgentID 代理商id
	AgentID uint64 `json:"agent_id,omitempty"`
	// StartDate 开始时间，格式 yyyy-MM-dd，最远可以查询3年内的数据
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，格式 yyyy-MM-dd，最远可以查询3年内的数据
	EndDate string `json:"end_date,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 每页条数
	PageSize int `json:"page_size,omitempty"`
	// Filtering  过滤器
	Filtering *TransferTransactionRecordFilter `json:"filtering,omitempty"`
}

// TransferTransactionRecordFilter 查询代理商转账记录过滤条件
type TransferTransactionRecordFilter struct {
	// CustomerIDs 转入/转出方客户
	CustomerIDs []uint64 `json:"customer_ids,omitempty"`
	// AccountIDs 转入/转出方账户
	AccountIDs []uint64 `json:"account_ids,omitempty"`
	// Payees 转入方账号
	Payees []uint64 `json:"payees,omitempty"`
	// PayeeCustomerIDs 转入方客户
	PayeeCustomerIDs []uint64 `json:"payee_customer_ids,omitempty"`
	// Remitters 转出方账户
	Remitters []uint64 `json:"remitters,omitempty"`
	// RemitterCustomerIDs 转出方客户
	RemitterCustomerIDs []uint64 `json:"remitter_customer_ids,omitempty"`
	// OperatorIDs 操作人
	OperatorIDs []uint64 `json:"operator_ids,omitempty"`
	// TransferOrderSerial 转账编号
	TransferOrderSerial string `json:"transfer_order_serial,omitempty"`
	// TransferType 转账类型 可选值:
	// ADD_MONEY 加款
	// MUTUAL_TRANSFER 同级账户转账
	// REFUND_MONEY 退款
	TransferType string `json:"transfer_type,omitempty"`
	// RemitterType 转出方账户类型 可选值:
	// ROLE_ADVERTISER 广告主
	// ROLE_AGENT 广告代理商
	// ROLE_CHILD_AGENT 代理子账户
	// ROLE_ECP_VIRTUAL_ADVERTISER 千川虚拟广告主
	// ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER 本地虚拟广告主
	// ROLE_VIRTAUL_ADVERTISER 虚拟广告主
	RemitterType enum.AdvertiserRole `json:"remitter_type,omitempty"`
	// PayeeType 转入方账户类型 可选值:
	// ROLE_ADVERTISER 广告主
	// ROLE_AGENT 广告代理商
	// ROLE_CHILD_AGENT 代理子账户
	// ROLE_ECP_VIRTUAL_ADVERTISER 千川虚拟广告主
	// ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER 本地虚拟广告主
	// ROLE_VIRTAUL_ADVERTISER 虚拟广告主
	PayeeType enum.AdvertiserRole `json:"payee_type,omitempty"`
	// Platform 业务平台 可选值:
	// AD 巨量广告
	// EFFECT 效果联盟
	// STAR 巨量星图
	// TRUSTEESHIP_MARKETING 托管营销
	// ZTT 智投通
	Platform string `json:"platform,omitempty"`
	// RemitterFirstAdAgentID 转出方代理商账户
	RemitterFirstAdAgentID uint64 `json:"remitter_first_ad_agent_id,omitempty"`
	// PayeeFirstAdAgentID 转入方代理商账户
	PayeeFirstAdAgentID uint64 `json:"payee_first_ad_agent_id,omitempty"`
}

// Encode implements GetRequest interfafce
func (r TransferTransactionRecordRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("agent_id", strconv.FormatUint(r.AgentID, 10))
	values.Set("start_date", r.StartDate)
	values.Set("end_date", r.EndDate)
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// TransferTransactionRecordResponse 查询代理商转账记录 API Response
type TransferTransactionRecordResponse struct {
	model.BaseResponse
	Data *TransferTransactionRecordResult `json:"data,omitempty"`
}

type TransferTransactionRecordResult struct {
	Records  []TransferTransactionRecord `json:"records,omitempty"`
	PageInfo *model.PageInfo             `json:"page_info,omitempty"`
}

// TransferTransactionRecord 代理商转账记录
type TransferTransactionRecord struct {
	// ModifyTime 转账时间
	ModifyTime string `json:"modity_time,omitempty"`
	// Remitter 转账时间转出方账户ID
	Remitter model.Uint64 `json:"remitter,omitempty"`
	// RemitterFirstAdAgentID 转出方一代账户ID
	RemitterFirstAdAgentID model.Uint64 `json:"remitter_first_ad_agent_id,omitempty"`
	// RemitterFirstAdAgentName 转出方一代账户名称
	RemitterFirstAdAgentName string `json:"remitter_first_ad_agent_name,omitempty"`
	// RemitterSecondAdAgentID 转出方二代账户ID
	RemitterSecondAdAgentID model.Uint64 `json:"remitter_second_ad_agent_id,omitempty"`
	// PayeeFirstAdAgentID 转入方一代账户ID
	PayeeFirstAdAgentID model.Uint64 `json:"payee_first_ad_agent_id,omitempty"`
	// PayeeFirstAdAgentName 转入方一代账户名称
	PayeeFirstAdAgentName string `json:"payee_first_ad_agent_name,omitempty"`
	// PayeeSecondAdAgentID 转入方二代账户ID
	PayeeSecondAdAgentID model.Uint64 `json:"payee_second_ad_agent_id,omitempty"`
	// RemitterName 转出方账户名称
	RemitterName string `json:"remitter_name,omitempty"`
	// RemitterCustomerName 转出方客户名称
	RemitterCustomerName string ` json:"remitter_customer_name,omitempty"`
	// RemitterCustomerID 转出方客户ID
	RemitterCustomerID model.Uint64 `json:"remitter_customer_id,omitempty"`
	// Payee 转入方账户ID
	Payee model.Uint64 `json:"payee,omitempty"`
	// PayeeName 转入方账户名称
	PayeeName string `json:"payee_name,omitempty"`
	// PayeeRole 转入方账户类型:
	// ROLE_ADVERTISER 普通广告主
	// ROLE_ADVERTISER_SYSTEM_ACCOUNT 广告主系统账号
	// ROLE_AGENT 广告代理商
	// ROLE_AGENT_SYSTEM_ACCOUNT 代理商系统账号
	// ROLE_CHILD_AGENT 代理子账户
	// ROLE_ECP_VIRTUAL_ADVERTISER 千川虚拟广告主
	// ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER 本地虚拟广告主
	// ROLE_VIRTAUL_ADVERTISER 虚拟广告主
	PayeeRole enum.AdvertiserRole `json:"payee_role,omitempty"`
	// RemitterRole 转出方账户类型 :
	// ROLE_ADVERTISER 普通广告主
	// ROLE_ADVERTISER_SYSTEM_ACCOUNT 广告主系统账号
	// ROLE_AGENT 广告代理商
	// ROLE_AGENT_SYSTEM_ACCOUNT 代理商系统账号
	// ROLE_CHILD_AGENT 代理子账户
	// ROLE_ECP_VIRTUAL_ADVERTISER 千川虚拟广告主
	// ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER 本地虚拟广告主
	// ROLE_VIRTAUL_ADVERTISER 虚拟广告主
	RemitterRole enum.AdvertiserRole `json:"remitter_role,omitempty"`
	// PayeeCustomerName 转入方客户名称
	PayeeCustomerName string `json:"payee_customer_name,omitempty"`
	// PayeeCustomerID 转入方客户ID
	PayeeCustomerID model.Uint64 `json:"payee_customer_id,omitempty"`
	// Amount 总金额（元）
	Amount float64 `json:"amount,omitempty"`
	// PrepayAmount 预付金额（元）
	PrepayAmount float64 `json:"prepay_amount,omitempty"`
	// CreditAmount 授信金额额（元）
	CreditAmount float64 `json:"credit_amount,omitempty"`
	// Cash 非赠款金额(元)
	Cash float64 `json:"cash,omitempty"`
	// Grants 赠款金额(元)
	Grants float64 `json:"grants,omitempty"`
	// TransferType 转账类型:
	// ADD_MONEY
	// MUTUAL_TRANSFER
	// REFUND_MONEY
	TransferType string `json:"transfer_type,omitempty"`
	// Platform 业务平台:
	// AD 巨量广告
	// EFFECT 效果联盟
	// STAR 巨量星图
	// TRUSTEESHIP_MARKETING 托管营销
	// ZTT 智投通
	Platform string `json:"platform,omitempty"`
	// Operator 操作人名称
	Operator string `json:"operator,omitempty"`
	// OperatorID 操作人ID
	OperatorID model.Uint64 `json:"operator_id,omitempty"`
	// TransferOrderSerial 转账订单号
	TransferOrderSerial string `json:"transfer_order_serial,omitempty"`
	// TransferTargetPayStatus 转账状态 可选值:
	// NO_TRANSFER 未转账
	// TRANSFER_FAILED 转账失败
	// TRANSFER_ING 转账中
	// TRANSFER_PART TRANSFER_PART
	// TRANSFER_SUCCESS 转账成功
	TransferTargetPayStatus enum.TransferStatus `json:"transfer_target_pay_status,omitempty"`
	// Remark 转账备注
	Remark string `json:"remark,omitempty"`
}
