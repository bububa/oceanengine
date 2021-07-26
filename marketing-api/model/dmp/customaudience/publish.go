package customaudience

import "encoding/json"

// PublishRequest 发布人群包
type PublishRequest struct {
	// AdvertiserID 人群包所属广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitepmty"`
	// CustomAudienceID 人群包ID
	CustomAudienceID uint64 `json:"custom_audience_id,omitempty"`
}

// Encode implement PostRequest interface
func (r PublishRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
