package customaudience

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// PushRequest 推送人群包
type PushRequest struct {
	// AdvertiserID 人群包所属广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CustomAudienceID 人群包ID
	CustomAudienceID uint64 `json:"custom_audience_id,omitempty"`
	// TargetAdvertiserIDs 推送广告主ID列表，最多推送100个广告主
	TargetAdvertiserIDs []uint64 `json:"target_advertiser_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r PushRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
