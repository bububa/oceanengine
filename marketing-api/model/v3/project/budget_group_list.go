package project

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// BudgetGroupListRequest 获取预算组列表 API Request
type BudgetGroupListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	Filtering *BudgetGroupListFilter `json:"filtering,omitempty"`
	// StatisticsTime 统计时间范围
	StatisticsTime *model.DateRange `json:"statistics_time,omitempty"`
	// Page 页数，默认值：1，page范围为[1,99999]
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值：10，page_size范围为[1,100]
	PageSize int `json:"page_size,omitempty"`
}

type BudgetGroupListFilter struct {
	// BudgetGroupIDs 预算组ID列表，一次最多传入200个ID
	BudgetGroupIDs []uint64 `json:"budget_group_ids,omitempty"`
	// BudgetGroupKeyword 预算组名称模糊搜索关键字
	BudgetGroupKeyword string `json:"budget_group_keyword,omitempty"`
	// DeliveryType 投放类型 可选值:
	// MANUAL 手动投放模式
	// PROCEDURAL 自动投放模式
	DeliveryType enum.DeliveryMode `json:"delivery_type,omitempty"`
	// AdType 广告类型 可选值:
	// ALL 通投广告
	// SEARCH 搜索广告
	AdType enum.CampaignType `json:"ad_type,omitempty"`
	// BudgetGroupStatusFirst 预算组聚合一级状态 可选值:
	// DELETED 已删除
	// ENABLE 启用中
	// UNDELIVERIED 未投放
	BudgetGroupStatusFirst enum.BudgetGroupStatus `json:"budget_group_status_first,omitempty"`
	// BudgetGroupStatusSecond 预算组聚合二级状态 可选值:
	// ACCOUNT_EXCEEDED 账户超出预算
	// GROUP_EXCEEDED 预算组超出预算
	BudgetGroupStatusSecond enum.BudgetGroupStatus `json:"budget_group_status_second,omitempty"`
}

// Encode implement GetRequest interface
func (r BudgetGroupListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.StatisticsTime != nil {
		values.Set("statistics_time", string(util.JSONMarshal(r.StatisticsTime)))
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

// BudgetGroupListResponse 获取预算组列表 API Response
type BudgetGroupListResponse struct {
	model.BaseResponse
	Data *BudgetGroupListResult `json:"data,omitempty"`
}

type BudgetGroupListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// BudgetGroups 预算组信息
	BudgetGroups []BudgetGroup `json:"budget_groups,omitempty"`
}

// BudgetGroup 预算组信息
type BudgetGroup struct {
	// BudgetGroupID 预算组ID
	BudgetGroupID uint64 `json:"budget_group_id,omitempty"`
	// BudgetGroupName 预算组名称
	BudgetGroupName string `json:"budget_group_name,omitempty"`
	// BudgetGroupBudget 预算组预算
	BudgetGroupBudget float64 `json:"budget_group_budget,omitempty"`
	// AdType 广告类型 可选值:
	// ALL 通投广告
	// SEARCH 搜索广告
	AdType enum.CampaignType `json:"ad_type,omitempty"`
	// DeliveryType 投放类型 可选值:
	// MANUAL 手动投放模式
	// PROCEDURAL 自动投放模式
	DeliveryType enum.DeliveryMode `json:"delivery_type,omitempty"`
	// BudgetGroupStatusFirst 预算组聚合一级状态 可选值:
	// DELETED 已删除
	// ENABLE 启用中
	// UNDELIVERIED 未投放
	BudgetGroupStatusFirst enum.BudgetGroupStatus `json:"budget_group_status_first,omitempty"`
	// BudgetGroupStatusSecond 预算组聚合二级状态 可选值:
	// ACCOUNT_EXCEEDED 账户超出预算
	// GROUP_EXCEEDED 预算组超出预算
	BudgetGroupStatusSecond enum.BudgetGroupStatus `json:"budget_group_status_second,omitempty"`
	// NumProjects 关联的（非删除）项目个数
	NumProjects int `json:"num_projects,omitempty"`
	// Cost 预算组消耗（元）
	Cost float64 `json:"cost,omitempty"`
}
