package advertiser

import (
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AccountBudgetUpdateRequest 更新账户日预算 API Request
type AccountBudgetUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// BudgetMode 预算类型
	// INFINITE 不限
	// SPECIFIED 日预算
	BudgetMode qianchuan.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 账户今日竞价预算，单位元
	Budget float64 `json:"budget,omitempty"`
}

func (r AccountBudgetUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
