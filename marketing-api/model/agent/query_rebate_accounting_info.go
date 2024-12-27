package agent

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QueryRebateAccountingInfoRequest 返点-查询返点核算流水 API Request
type QueryRebateAccountingInfoRequest struct {
	// AgentID 代理商
	AgentID uint64 `json:"agent_id,omitempty"`
	// Filtering 过滤器
	Filtering *QueryRebateAccountingInfoFilter `json:"filtering,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	PageSize int `json:"page_size,omitempty"`
}

type QueryRebateAccountingInfoFilter struct {
	// ApprovalStatus 返点核算单审批状态 ， 可选值:
	// AGENT_APPROVING 待代理商审批
	// BUSINESS_APPROVING 待平台复核终审
	// APPROVED 审批完成
	ApprovalStatus enum.RebateApprovalStatus `json:"approval_status,omitempty"`
	// Year 结算周期年 例:2024
	Year int `json:"year,omitempty"`
	// MonthQuarter 结算季度／月，可选值
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
	// IsCreateRebate 是否已创建返点 ，可选值:
	// YES 已生成返点
	// NO 未生成返点
	IsCreateRebate enum.YesNo `json:"is_create_rebate,omitempty"`
	// RebateBalanceSerial 返点流水号
	RebateBalanceSerial string `json:"rebate_balance_serial,omitempty"`
	// ContractSerial 合同编号
	ContractSerial string `json:"contract_serial,omitempty"`
	// RebateAccountingInfoSerial 返点核算信息编号
	RebateAccountingInfoSerial string `json:"rebate_accounting_info_serial,omitempty"`
	// UseType 使用类型，可选值
	// CASH 现金
	// CHARGE 充值
	// HEDGING 抵扣
	// NONPAYMENT 无需支付
	// REVERT 负数冲抵
	UseType enum.RebateUseType `json:"use_type,omitempty"`
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
func (r QueryRebateAccountingInfoRequest) Encode() string {
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

// QueryRebateAccountingInfoResponse 返点-查询返点核算流水 API Response
type QueryRebateAccountingInfoResponse struct {
	model.BaseResponse
	Data *QueryRebateAccountingInfoResult `json:"data,omitempty"`
}

type QueryRebateAccountingInfoResult struct {
	// PageInfo 翻页信息
	PageInfo *model.PageInfo        `json:"page_info,omitempty"`
	InfoList []RebateAccountingInfo `json:"info_list,omitempty"`
}

// RebateAccountingInfo 返点核算流水
type RebateAccountingInfo struct {
	// RebateAccountingInfoID 返点核算信息ID
	RebateAccountingInfoID uint64 `json:"rebate_accounting_info_id,omitempty"`
	// RebateAccountingInfoSerial 返点核算信息编号（与平台披露编号一致，建议使用）
	RebateAccountingInfoSerial string `json:"rebate_accounting_info_serial,omitempty"`
	// Status 审批状态code
	Status int `json:"status,omitempty"`
	// StatusName 审批状态namteme
	StatusName string `json:"status_name,omitempty"`
	// IsCreateRebate 是否已创建返点 1:是 2:否
	IsCreateRebate int `json:"is_create_rebate,omitempty"`
	// IsCreateRebateName 是否已创建返点name ("是"、"否")
	IsCreateRebateName string `json:"is_create_rebate_name,omitempty"`
	// RebateBalanceID 返点流水ID
	RebateBalanceID uint64 `json:"rebate_balance_id,omitempty"`
	// RebateBalanceSerial 返点流水编号（与平台披露编号一致，建议使用）
	RebateBalanceSerial string `json:"rebate_balance_serial,omitempty"`
	// UseTypes 使用类型list        se
	UseTypes []int `json:"use_types,omitempty"`
	// UseTypeNames 使用类型name
	UseTypeNames []string `json:"use_type_names,omitempty"`
	// Platform 业务线/平台
	Platform int `json:"platform,omitempty"`
	// PlatformName 业务线/平台name
	PlatformName string `json:"platform_name,omitempty"`
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
	// Amount 返点金额 单位 元 保留2位小数
	Amount float64 `json:"amount,omitempty"`
	// PerformanceAmount 业绩基数 单位 元 保留2位小数
	PerformanceAmount float64 `json:"performanctempe_amount,omitempty"`
	// Type 返点结算类 in型
	Type int `json:"type,omitempty"`
	// TypeName 返点结算类型namNae
	TypeName string `json:"type_name,omitempty"`
	// ContractSerial 合同编号
	ContractSerial string `json:"contract_serial,omitempty"`
	// ContractSubjectName 媒体签约主体
	ContractSubjectName string `json:"contract_subject_name,omitempty"`
	// CreateTime 创建时间 2000-01-01 00:00:00
	CreateTime string `json:"create_time,omitempty"`
	// CurrentApprovalStatusName 当前审批节点名称 （待代理商审批/待平台复核终审/审批完成）
	CurrentApprovalStatusName string `json:"current_approval_status_name,omitempty"`
}
