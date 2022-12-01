package project

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// BudgetUpdateRequest 更新项目预算 API Request
type BudgetUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Data 批量修改，包含广告id和出价
	Data []BudgetUpdateData `json:"data,omitempty"`
}

// BudgetUpdateData 修改信息
type BudgetUpdateData struct {
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// BudgetMode 允许值:
	// BUDGET_MODE_DAY日预算
	// BUDGET_MODE_INFINITE不限
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 预算，单位“元”，精度：两位小数。
	Budget float64 `json:"budget,omitempty"`
}

// Encode implement PostRequest interface
func (r BudgetUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
