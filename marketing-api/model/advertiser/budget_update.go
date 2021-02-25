package advertiser

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
)

type UpdateBudgetRequest struct {
	AdvertiserID uint64          `json:"advertiser_id,omitempty"`
	BudgetMode   enum.BudgetMode `json:"budget_mode,omitempty"`
	Budget       float64         `json:"budget,omitempty"`
}

func (r UpdateBudgetRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
