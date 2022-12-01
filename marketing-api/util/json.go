package util

import (
	"encoding/json"
)

// JSONMarshal encode json without html escape
func JSONMarshal(req interface{}) []byte {
	bs, _ := json.Marshal(req)
	return bs
}
