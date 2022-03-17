package advertiser

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateBudgetRequest 更新账户日预算 API Request
type UpdateBudgetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// BudgetMode 预算模式
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 预算值，取值范围：1000～9999999.99，当budget_mode=BUDGET_MODE_DAY时，必填
	// 单位：元； 最小值：0； 精度：小数点后两位
	Budget float64 `json:"budget,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateBudgetRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
