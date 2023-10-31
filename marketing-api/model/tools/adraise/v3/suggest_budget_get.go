package v3

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SuggestBudgetGetRequest 获取广告建议起量预算 API Request
type SuggestBudgetGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionIDs 广告ID列表，目前仅支持1个
	PromotionIDs []uint64 `json:"promotion_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r SuggestBudgetGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if len(r.PromotionIDs) > 0 {
		values.Set("delivery_range", string(util.JSONMarshal(r.PromotionIDs)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret

}

// SuggestBudgetGetResponse 获取广告建议起量预算 API Response
type SuggestBudgetGetResponse struct {
	model.BaseResponse
	Data struct {
		// List 预算列表
		List []SuggestBudget `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// SuggestBudget 预算
type SuggestBudget struct {
	// PromotionID 广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// PromotionBudget 广告预算
	PromotionBudget float64 `json:"promotion_budget,omitempty"`
	// SuggestBudget 建议起量预算
	SuggestBudget float64 `json:"suggest_budget,omitempty"`
}
