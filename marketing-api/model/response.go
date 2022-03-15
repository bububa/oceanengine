package model

import (
	"strconv"
	"strings"
)

// Response api response interface
type Response interface {
	// IsError 是否返回错误
	IsError() bool
	// Error implement error interface
	Error() string
}

// BaseResponse shared api response data fields
type BaseResponse struct {
	// Code 返回码
	Code int `json:"code"`
	// Message 返回信息
	Message string `json:"message"`
	// RequestID 请求的日志id，唯一标识一个请求
	RequestID string `json:"request_id,omitempty"`
}

// IsError implement Response interface
func (r BaseResponse) IsError() bool {
	return r.Code != 0
}

// Error implement Response interface
func (r BaseResponse) Error() string {
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(r.Code))
	builder.WriteString(":")
	builder.WriteString(r.Message)
	return builder.String()
}
