package creative

// Error 错误信息
type Error struct {
	// Code 返回码
	Code int `json:"code,omitempty"`
	// Message 返回信息
	Message string `json:"message,omitempty"`
}
