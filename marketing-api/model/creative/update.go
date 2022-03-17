package creative

import (
	"encoding/json"
)

// UpdateRequest 修改创意信息 API Request
type UpdateRequest struct {
	CreativeDetail
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
