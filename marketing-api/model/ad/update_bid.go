package ad

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

type UpdateBidRequest struct {
	AdvertiserID uint64                 `json:"advertiser_id,omitempty"` // 广告主ID
	Data         []UpdateBidRequestList `json:"data,omitempty"`
}

type UpdateBidRequestList struct {
	AdID uint64  `json:"ad_id,omitempty"`
	Bid  float64 `json:"bid,omitempty"`
}

func (r UpdateBidRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

type UpdateBidResponse struct {
	model.BaseResponse
	Data *UpdateBidResponseData `json:"data,omitempty"`
}

type UpdateBidResponseData struct {
	AdIDs []uint64 `json:"ad_ids,omitempty"`
}
