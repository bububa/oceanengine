package project

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// BudgetGroupUpdateRequest 更新预算组 API Request
type BudgetGroupUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// BudgetGroupID 需要更新的预算组ID
	BudgetGroupID uint64 `json:"budget_group_id,omitempty"`
	// BudgetGroupName 预算组名称，50个字以内，预算组之间名称不可重复，不可为空
	BudgetGroupName string `json:"budget_group_name,omitempty"`
	// BudgetGroupBudget 日预算，单位：元，范围：[300，9,999,999]，预算组中的每个项目在实际消耗时，以【项目、预算组】中较低预算为准
	BudgetGroupBudget float64 `json:"budget_group_budget,omitempty"`
}

// Encode implement PostRequest interface
func (r BudgetGroupUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// BudgetGroupUpdateResponse 更新预算组 API Response
type BudgetGroupUpdateResponse struct {
	model.BaseResponse
	Data struct {
		// BudgetGroupID 预算组ID
		BudgetGroupID uint64 `json:"budget_group_id,omitempty"`
	} `json:"data,omitempty"`
}
