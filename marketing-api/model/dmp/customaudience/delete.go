package customaudience

import "encoding/json"

// DeleteRequest 删除人群包
type DeleteRequest struct {
	// AdvertiserID 人群包所属广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitepmty"`
	// CustomAudienceID 人群包ID
	CustomAudienceID uint64 `json:"custom_audience_id,omitempty"`
}

// Encode implement PostRequest interface
func (r DeleteRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
