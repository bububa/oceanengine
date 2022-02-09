package creativecomponent

import (
	"encoding/json"
)

// UpdateRequest 更新组件 API Request
type UpdateRequest struct {
	ComponentID uint64 `json:"component_id,omitempty"`
	PostRequest
}

func (r UpdateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)

	return ret
}
