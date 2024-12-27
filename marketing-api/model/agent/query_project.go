package agent

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QueryProjectRequest 查询项目 API Request
type QueryProjectRequest struct {
	// AgentID 代理商ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// Filtering 过滤器
	Filtering *QueryProjectFilter `json:"filtering,omitempty"`
	// Count 页码游标值（单页限制，最大1000，必填）
	Count int `json:"count,omitempty"`
	// Cursor 页码游标值（根据cursor进行滚动查询，第一次传-1，后续传前一次返回列表中的最小值）
	Cursor int `json:"cursor,omitempty"`
}

type QueryProjectFilter struct {
	// CustomerID 客户ID
	CustomerID uint64 `json:"customer_id,omitempty"`
	// PlatformList 平台列表，可选值
	// AD:巨量广告
	// NATIVELIFE:本地推
	// QIANCHUAN:千川
	PlatformList []string `json:"platform_list,omitempty"`
	// ServingTypeList 投放类型
	// PINPAI：品牌
	// PINPAIPMP：品牌PMP
	// XIAOGUO：效果
	// XIAOGUO_RTB：效果RTB
	ServingTypeList []string `json:"serving_type_list,omitempty"`
	// ProjectStatusList 项目状态
	// APPROVED：审批通过
	ProjectStatusList []string `json:"project_status_list,omitempty"`
	// ProjectStartDateBegin 项目开始时间-开始 例如 2023-07-06
	ProjectStartDateBegin string `json:"project_start_date_begin,omitempty"`
	// ProjectStartDateEnd 项目开始时间-结束 例如 2023-07-06
	ProjectStartDateEnd string `json:"project_start_date_end,omitempty"`
	// ProjectEndDateBegin 项目结束时间-开始 例如 2023-07-06
	ProjectEndDateBegin string `json:"project_end_date_begin,omitempty"`
	// ProjectEndDateEnd 项目结束时间-结束 例如 2023-07-06
	ProjectEndDateEnd string `json:"project_end_date_end,omitempty"`
	// ReceiptStatusList 项目回款状态
	// COMPLETED：已回款
	// PART_COMPLETED：部分回款
	// UNCOMPLETED：未回款
	ReceiptStatusList []string `json:"receipt_status_list,omitempty"`
	// Deadline 应回款日期
	Deadline string `json:"deadline,omitempty"`
	// AdvertiserID 投放账号ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implements GetRequest interface
func (r QueryProjectRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("agent_id", strconv.FormatUint(r.AgentID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	values.Set("cursor", strconv.Itoa(r.Cursor))
	values.Set("count", strconv.Itoa(r.Count))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QueryProjectResponse 查询项目信息 API Response
type QueryProjectResponse struct {
	model.BaseResponse
	Data *QueryProjectResult `json:"data,omitempty"`
}

type QueryProjectResult struct {
	// ProjectInfoList 项目列表
	ProjectInfoList []ProjectInfo `json:"project_info_list,omitempty"`
	// Total 总量
	Total int `json:"total,omitempty"`
	// Cursor 页码游标值
	Cursor int `json:"cursor,omitempty"`
}

// ProjectInfo 项目
type ProjectInfo struct {
	// ID 项目ID
	ID uint64 `json:"id,omitempty"`
	// Serial 项目编号
	Serial string `json:"serial,omitempty"`
	// Name 项目名称
	Name string `json:"name,omitempty"`
	// PayMethodName 付款方式
	PayMethodName string `json:"pay_method_name,omitempty"`
	// PojectStartDate 项目开始时间 例如 2023-07-06                                   r
	ProjectStartDate string `json:"project_start_date,omitempty"`
	// ProjectEndDate 项目结束时间 例如 2023-07-06
	ProjectEndDate string `json:"project_end_date,omitempty"`
	// ProjectStatus 项目状态
	ProjectStatus int `json:"project_status,omitempty"`
	// PriojectStatusName 项目状态名称
	ProjectStatusName string `json:"project_status_name,omitempty"`
	// ConfirmStatus 确认状t态
	ConfirmStatus int `json:"confirm_status,omitempty"`
	// ConfirmStatusName 确认状态名称
	ConfirmStatusName string `json:"confirm_status_name,omitempty"`
	// ProjectSubType 项目类型
	ProjectSubType int `json:"project_sub_type,omitempty"`
	// ProjectSubTypeName 项目类型名称
	ProjectSubTypeName string `json:"project_sub_type_name,omitempty"`
	// ServingType 投放方式
	ServingType int `json:"serving_type,omitempty"`
	// ServingTypeName 投放方式名称
	ServingTypeName string `json:"serving_type_name,omitempty"`
	// Amount 项目金额（元）
	Amount model.Float64 `json:"amount,omitempty"`
	// ReceiptStatus 回款状态
	ReceiptStatus int `json:"receipt_status,omitempty"`
	// ReceiptStatusName 回款状态名称
	ReceiptStatusName string `json:"receipt_status_name,omitempty"`
	// Platform 平台
	Platform int `json:"platform,omitempty"`
	// PlatformName 平台名称
	PlatformName string `json:"platform_name,omitempty"`
	// CustomerID 客户ID
	CustomerID uint64 `json:"customer_id,omitempty"`
	// CustomerName 客户名称
	CustomerName string `json:"customer_name,omitempty"`
	// AgentID 代理商ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// AgentName 代理商名称
	AgentName string `json:"agent_name,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdvertiserName 广告主名称
	AdvertiserName string `json:"advertiser_name,omitempty"`
	// AdcSerial 广告发布协议编号
	AdcSerial string `json:"adc_serial,omitempty"`
	// ContractSerial 主合同编号
	ContractSerial string `json:"contract_serial,omitempty"`
	// ReceiptSubjectName 我方主体名称
	ReceiptSubjectName string `json:"receipt_subject_name,omitempty"`
	// Deadline 应回款时间
	Deadline string `json:"deadline,omitempty"`
	// RemainVerificationAmount 待回款金额（元）
	RemainVerificationAmount model.Float64 `json:"remain_verification_amount,omitempty"`
	// RechargeCost 现金消耗
	RechargeCost model.Float64 `json:"recharge_cost,omitempty"`
	// PreepayCost 预付消耗
	PrepayCost model.Float64 `json:"prepay_cost,omitempty"`
	// CreditCost 授信消耗
	CreditCost model.Float64 `json:"credit_cost,omitempty"`
	// GrantCost 赠款消耗
	GrantCost model.Float64 `json:"grant_cost,omitempty"`
	// InvoiceAmount 已开票金额
	InvoiceAmount model.Float64 `json:"invoice_amount,omitempty"`
}
