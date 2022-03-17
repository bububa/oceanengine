package ad

import "github.com/bububa/oceanengine/marketing-api/util"

// UpdateBudgetRequest 更新计划预算 API Request
type UpdateBudgetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Data 批量修改，包含计划id和出价
	Data []UpdateBudgetRequestList `json:"data,omitempty"`
}

// UpdateBudgetRequestList 修改信息
type UpdateBudgetRequestList struct {
	// AdID 广告计划ID，广告计划ID需要属于广告主
	AdID uint64 `json:"ad_id,omitempty"`
	// Budget 预算，单位“元”，精度：两位小数。
	Budget float64 `json:"budget,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateBudgetRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
