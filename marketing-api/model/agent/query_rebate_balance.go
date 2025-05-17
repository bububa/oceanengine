package agent

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QueryRebateBalanceRequest 返点-查询返点流水 API Request
type QueryRebateBalanceRequest struct {
	// AgentID 代理商ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// Filtering 过滤器
	Filtering *QueryRebateBalanceFilter `json:"filtering,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	PageSize int `json:"page_size,omitempty"`
}

type QueryRebateBalanceFilter struct {
	// RebateBalanceSerial 返点流水编号 （与平台披露编号一致，建议使用）
	RebateBalanceSerial string `json:"rebate_balance_serial,omitempty"`
	// UseType 使用类型 ，可选值:
	// CASH 现金
	// CHARGE 充值
	// HEDGING 抵扣
	// NONPAYMENT 无需支付
	// REVERT 负数冲抵
	UseType enum.RebateUseType `json:"use_type,omitempty"`
	// ApprovalStatus 返点流水审核状态，可选值:
	// AGENT_APPROVING 待代理商审批
	// APPROVED 审批完成
	ApprovalStatus enum.RebateApprovalStatus `json:"approval_status,omitempty"`
	// Year 结算周期年
	Year int `json:"year,omitempty"`
	// MonthQuarter 所属季度／月 ， 可选值:
	// AllYear 全年
	// Apr 4月
	// Aug 8月
	// Dec 12月
	// Feb 2月
	// FirHalf 上半年
	// Jan 1月
	// Jul 7月
	// Jun 6月
	// Mar 3月
	// May 5月
	// Nov 11月
	// Oct 10月
	// Q1 第一季度
	// Q2 第二季度
	// Q3 第三季度
	// Q4 第四季度
	// SecHalf 下半年
	// Sep 9月
	MonthQuarter string `json:"month_quarter,omitempty"`
	// ContractSerial 合同编号
	ContractSerial string `json:"contract_serial,omitempty"`
	// CreateStartDate 创建时间范围开始时间
	CreateStartDate string `json:"create_start_date,omitempty"`
	// CreateEndDate 创建时间范围结束时间
	CreateEndDate string `json:"create_end_date,omitempty"`
	// InvoiceStatusList 开票状态， 可选值:
	// NO 未开票
	// PART 部分开票
	// COMPLETE 已开票
	InvoiceStatusList []enum.InvoiceStatus `json:"invoice_status_list,omitempty"`
	// StampStatusList 盖章状态， 可选值:
	// NO 未盖章
	// STAMPING 审批中
	// STAMPED 已盖章
	StampStatusList []enum.StampStatus `json:"stamp_status_list,omitempty"`
	// ArchiveStatusList 回收状态， 可选值:
	// NO_RETRIEVE 未回收
	// RETRIEVED 已回收
	// ARCHIVED 已归档
	ArchiveStatusList []enum.ArchiveStatus `json:"archive_status_list,omitempty"`
	// Standard 是否标准， 可选值:
	// YES 标准
	// NO 非标
	// YES_OF_INSPIRE 标准-激励
	Standard string `json:"standard,omitempty"`
	// StampType 是否电子签章 ， 可选值:
	// ONLINE 电子签章
	// DEFAULT 线下签章
	StampType enum.StampType `json:"stamp_type,omitempty"`
	// Platforms 业务线/平台，可选值
	// AD 巨量广告
	// ECOMMERCE 电商平台
	// FC 房产
	// NATIVE 巨量本地推
	// QC 千川
	// STAR 星图
	Platforms []string `json:"platforms,omitempty"`
}

// Encode implements GetRequest interface
func (r QueryRebateBalanceRequest) Encode() string {
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

// QueryRebateBalanceResponse 返点-查询返点流水 API Response
type QueryRebateBalanceResponse struct {
	model.BaseResponse
	Data *QueryRebateBalanceResult `json:"data,omitempty"`
}

type QueryRebateBalanceResult struct {
	// PageInfo 翻页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	Rebates  []RebateBalance `json:"rebates,omitempty"`
}

type RebateBalance struct {
	// RebateBalanceID 返点流水ID
	RebateBalanceID uint64 `json:"rebate_balance_id,omitempty"`
	// RebateBalanceSerial 返点流水编号 （与平台披露编号一致，建议使用）
	RebateBalanceSerial string `json:"rebate_balance_serial,omitempty"`
	// Status 审批状态code
	Status int `json:"status,omitempty"`
	// StatusName 审批状态namteme
	StatusName string `json:"status_name,omitempty"`
	// UseTypes 使用类型list        se
	UseTypes []int `json:"use_types,omitempty"`
	// UseTypeNames 使用类型name
	UseTypeNames []string `json:"use_type_names,omitempty"`
	// AgentID 代理商ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// AgentName 代理商name
	AgentName string `json:"agent_name,omitempty"`
	// Year 结算年份 例：2024
	Year int `json:"year,omitempty"`
	// MonthQuarter 结算季度/月 （1:1月，2:2月，3:3月，4:4月，5:5月，6:6月，7:7月，8:8月，9:9月，10:10月，11:11月，12:12月，13:Q1，14:Q2，15:Q3，16:Q4，17:上半年，18:下半年，19:全年）
	MonthQuarter int `json:"month_quarter,omitempty"`
	// MonthQuarterName 结算季度/月name
	MonthhQuarterName string `json:"month_quarter_name,omitempty"`
	// ContractSerial 合同编号
	ContractSerial string `json:"contract_serial,omitempty"`
	// ContractName 合同名称
	ContractName string `json:"contract_name,omitempty"`
	// ContractSubjectName 媒体签约主体
	ContractSubjectName string `json:"contract_subject_name,omitempty"`
	// Amount 返点金额 单位 元 保留2位小数
	Amount float64 `json:"amount,omitempty"`
	// UsedAmount 已使用金额 单位 元 保留2位小数
	UsedAmount float64 `json:"used_amount,omitempty"`
	// FrozenAmount 冻结金额 单位 元 保留2位小数
	FrozenAmount float64 `json:"frozen_amount,omitempty"`
	// CreateTime 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// InvoiceStatus 开票状态
	InvoiceStatus int `json:"invoice_status,omitempty"`
	// InvoiceStatusName 开票状态名称
	InvoiceStatusName string `json:"invoice_status_name,omitempty"`
	// InvoiceAmount 实际开票金额 单位 元 保留2位小数
	InvoiceAmount float64 `json:"invoice_amount,omitempty"`
	// InvoiceFrozenAmount 已申请开票金额 单位 元 保留2位小数
	InvoiceFrozenAmount float64 `json:"invoice_frozen_amount,omitempty"`
	// InvoiceAmountRemain 可开票金额 单位 元 保留2位小数
	InvoiceAmountRemain float64 `json:"invoice_amount_remain,omitempty"`
	// StampStatus 盖章状态
	StampStatus int `json:"stamp_status,omitempty"`
	// StampStatusName 盖章状态名称
	StampStatusName string `json:"stamp_status_name,omitempty"`
	// ArchiveStatus 回收状态
	ArchiveStatus int `json:"archive_status,omitempty"`
	// ArchiveStatusName 回收状态名称
	ArchiveStatusName string `json:"archive_status_name,omitempty"`
	// Standard 是否标准
	Standard int `json:"standard,omitempty"`
	// StandardName 是否标准名称
	StandardName string `json:"standard_name,omitempty"`
	// IsOnlineStamp 签章类型名称
	IsOnlineStamp string `json:"is_online_stamp,omitempty"`
	// Platform 平台/业务线
	Platform int `json:"platform,omitempty"`
	// PlatformName 平台/业务线名称
	PlatformName string `json:"platform_name,omitempty"`
	// CurrentApprovalStatusName 当前审批节点名称
	CurrentApprovalStatusName string `json:"current_approval_status_name,omitempty"`
}
