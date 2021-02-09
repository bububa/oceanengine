package ad

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

type UpdateBudgetRequest struct {
	AdvertiserID uint64                    `json:"advertiser_id,omitempty"` // 广告主ID
	Data         []UpdateBudgetRequestList `json:"data,omitempty"`
}

type UpdateBudgetRequestList struct {
	AdID   uint64  `json:"ad_id,omitempty"`
	Budget float64 `json:"budget,omitempty"`
}

func (r UpdateBudgetRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

type UpdateBudgetResponse struct {
	model.BaseResponse
	Data *UpdateBudgetResponseData `json:"data,omitempty"`
}

type UpdateBudgetResponseData struct {
	AdIDs []uint64 `json:"ad_ids,omitempty"`
}
