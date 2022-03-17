package adconvert

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// PushRequest 转化目标推送 API Request
type PushRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ConvertID 转化目标ID
	ConvertID uint64 `json:"convert_id,omitempty"`
	// TargetAdvertiserIDs 推送的广告主ID列表，上限：50
	TargetAdvertiserIDs []uint64 `json:"target_advertiser_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r PushRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
