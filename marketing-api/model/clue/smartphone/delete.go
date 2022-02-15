package smartphone

import (
	"encoding/json"
)

// DeleteRequest 删除智能电话 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// InstanceID 智能电话ID
	InstanceID uint64 `json:"instance_id,omitempty"`
}

// Encode implenent PostRequest interface
func (r DeleteRequest) Encode() []byte {
	b, _ := json.Marshal(r)
	return b
}
