package ad

import (
	"encoding/json"
)

// UpdateBudgetRequest 更新计划预算 API Request
type UpdateBudgetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Data 批量修改预算，包含计划ID和预算，list长度限制1～100.
	Data []UpdateBudgetRequestList `json:"data,omitempty"`
}

// UpdateBudgetRequestList 批量修改预算
type UpdateBudgetRequestList struct {
	// AdID 广告计划ID，广告计划id需要属于广告主，且ad_id不能重复，否则会报错！
	AdID uint64 `json:"ad_id,omitempty"`
	// Budget 预算，单位：元。
	// 24小时内修改预算操作，不能超过20次，24小时是指自然天的24小时；
	// 单次修改预算幅度不能低于100元（增加或者减少）;
	// 修改后预算金额，不能低于当前已消费金额的105%，以整百单位向上取整；
	Budget float64 `json:"budget,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateBudgetRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
