package track

import (
	"strconv"
	"strings"
)

// Response 线索-API上报数据 API Response
type Response struct {
	// Code 返回值
	Code int `json:"code,omitempty"`
	// Msg 错误信息
	Msg string `json:"msg,omitempty"`
	// Ret .
	Ret int `json:"ret,omitempty"`
}

// IsError 是否为error
func (r Response) IsError() bool {
	return r.Code != 0
}

// Error implement error interface
func (r Response) Error() string {
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(r.Code))
	builder.WriteString(":")
	builder.WriteString(r.Msg)
	return builder.String()
}
