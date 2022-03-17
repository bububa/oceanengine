package customaudience

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeleteRequest 删除人群包
type DeleteRequest struct {
	// AdvertiserID 人群包所属广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CustomAudienceID 人群包ID
	CustomAudienceID uint64 `json:"custom_audience_id,omitempty"`
}

// Encode implement PostRequest interface
func (r DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
