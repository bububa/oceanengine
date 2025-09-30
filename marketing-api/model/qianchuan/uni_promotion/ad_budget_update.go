package unipromotion

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdBudgetUpdateRequest 更新全域推广计划预算 API Request
type AdBudgetUpdateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UpdateBudgetInfos 更新预算的计划id和预算价格列表
	// 注意：单次最多支持10个
	UpdateBudgetInfos []UpdateBudgetInfo `json:"update_budget_infos,omitempty"`
}

// UpdateBudgetInfo 更新预算的计划id和预算价格
type UpdateBudgetInfo struct {
	// AdID 商品全域计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// Budget 更新后的预算，单位元，最多支持两位小数
	Budget float64 `json:"budget,omitempty"`
	// MinEstimateConvert 预估最小转化
	// 注意：
	// 可通过【获取全域建议预算】接口获取预估值， 如果需要生效成本保障，需要传入正确的预估建议。
	// 对于商品全域，如果传入的预算为建议预算，该字段必填
	MinEstimateConvert int64 `json:"min_estimate_convert,omitempty"`
	// EstimateConvert 预估转化
	// 注意：
	// 可通过【获取全域建议预算】接口获取预估值， 如果需要生效成本保障，需要传入正确的预估建议。
	// 对于商品全域，如果传入的预算为建议预算，该字段必填
	EstimateConvert int64 `json:"estimate_convert,omitempty"`
	// EstimateRoiGoal 预估roi
	// 注意：
	// 可通过【获取全域建议预算】接口获取预估值， 如果需要生效成本保障，需要传入正确的预估建议。
	// 对于商品全域，如果传入的预算为建议预算，该字段必填
	EstimateRoiGoal float64 `json:"estimate_roi_goal,omitempty"`
	// MinEstimateRoiGoal 预估最小roi
	// 注意：
	// 可通过【获取全域建议预算】接口获取预估值， 如果需要生效成本保障，需要传入正确的预估建议。
	// 对于商品全域，如果传入的预算为建议预算，该字段必填
	MinEstimateRoiGoal float64 `json:"min_estimate_roi_goal,omitempty"`
}

// Encode implements PostRequest inteface
func (r AdBudgetUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AdUpdateResponse 更新全域推广计划 API Response
type AdUpdateResponse struct {
	model.BaseResponse
	Data struct {
		// Results 批量更新结果
		Results []AdUpdateResult `json:"results,omitempty"`
	} `json:"data"`
}

// AdUpdateResult 批量更新结果
type AdUpdateResult struct {
	// AdID 操作更新的计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// Status 操作状态，可选值:
	// SUCCESS更新成功
	// FAILED更新失败
	Status string `json:"status,omitempty"`
	// ErrorMessage 更新失败的原因
	ErrorMessage string `json:"error_message,omitempty"`
}

func (r AdUpdateResult) IsError() bool {
	return r.Status == "FAILED"
}

func (r AdUpdateResult) Error() string {
	return r.ErrorMessage
}
