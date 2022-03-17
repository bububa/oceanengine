package util

import (
	"bytes"
	"encoding/json"
)

// JSONMarshal encode json without html escape
func JSONMarshal(req interface{}) []byte {
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	encoder.Encode(req)
	return buf.Bytes()
}
