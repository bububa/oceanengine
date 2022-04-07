package spi

// BaseResponse 返回信息基类
type BaseResponse struct {
	// StatusCode 执行状态码，200表示成功，4xx 业务异常不重试，5xx 系统异常重试
	StatusCode int `json:"StatusCode,omitempty"`
	// StatusMessage 执行结果
	StatusMessage string `json:"StatusMessage,omitempty"`
}
