package project

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// BudgetGroupCreateRequest 创建预算组 API Request
type BudgetGroupCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// BudgetGroupName 预算组名称，50个字以内，预算组之间名称不可重复，不可为空
	BudgetGroupName string `json:"budget_group_name,omitempty"`
	// BudgetGroupBudget 日预算，单位：元，范围：[300，9,999,999]，预算组中的每个项目在实际消耗时，以【项目、预算组】中较低预算为准
	BudgetGroupBudget float64 `json:"budget_group_budget,omitempty"`
	// CampaignType 广告类型 可选值:
	// ALL 通投广告
	// SEARCH 搜索广告
	CampaignType enum.CampaignType `json:"campaign_type,omitempty"`
	// DeliveryMode 投放类型 可选值:
	// MANUAL 手动投放模式
	// PROCEDURAL 自动投放模式
	DeliveryMode enum.DeliveryMode `json:"delivery_mode,omitempty"`
}

// Encode implement PostRequest interface
func (r BudgetGroupCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// BudgetGroupCreateResponse 创建预算组 API Response
type BudgetGroupCreateResponse struct {
	model.BaseResponse
	Data struct {
		// BudgetGroupID 预算组ID
		BudgetGroupID uint64 `json:"budget_group_id,omitempty"`
	} `json:"data,omitempty"`
}
