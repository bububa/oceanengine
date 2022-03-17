package ad

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateBudgetRequest 更新计划预算 API Request
type UpdateBudgetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Data 更新预算的计划id和预算价格列表，最多支持10个
	Data []UpdateBudgetRequestData `json:"data,omitempty"`
}

// UpdateBudgetRequestData 更新预算的计划id和预算价格
type UpdateBudgetRequestData struct {
	// AdID 广告计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// Budget 更新后的预算，最多只有两位小数
	Budget float64 `json:"budget,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateBudgetRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
