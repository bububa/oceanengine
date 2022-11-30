package v3

import "github.com/bububa/oceanengine/marketing-api/util"

// BudgetUpdateRequest 更新计划预算 API Request
type BudgetUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Data 批量修改，包含计划id和出价
	Data []BudgetUpdateData `json:"data,omitempty"`
}

// BudgetUpdateData 修改信息
type BudgetUpdateData struct {
	// PromotionID 广告计划ID，广告计划ID需要属于广告主
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// Budget 预算，单位“元”，精度：两位小数。
	Budget float64 `json:"budget,omitempty"`
}

// Encode implement PostRequest interface
func (r BudgetUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
