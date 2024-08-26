package track

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/util"
)

// Response 线索-API上报数据 API Response
type Response struct {
	// Msg 错误信息
	Msg string `json:"msg,omitempty"`
	// RequestID 请求的日志id，唯一标识一个请求
	RequestID string `json:"request_id,omitempty"`
	// Code 返回值
	Code int `json:"code,omitempty"`
	// Ret .
	Ret int `json:"ret,omitempty"`
}

// IsError 是否为error
func (r Response) IsError() bool {
	return r.Code != 0
}

// Error implement error interface
func (r Response) Error() string {
	return util.StringsJoin(strconv.Itoa(r.Code), ":", r.Msg)
}

// APIRequestID implement Response interface
func (r Response) APIRequestID() string {
	return r.RequestID
}
