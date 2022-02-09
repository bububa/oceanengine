package creativecomponent

import (
	"encoding/json"
)

// UpdateRequest 更新组件 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ComponentID 组件ID
	ComponentID uint64 `json:"component_id,omitempty"`
	// ComponentInfo 组件信息
	ComponentInfo *ComponentInfo `json:"component_info,omitempty"`
}

// Encode implement Post Request interface
func (r UpdateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)

	return ret
}
