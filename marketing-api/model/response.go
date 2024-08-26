package model

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/util"
)

// Response api response interface
type Response interface {
	// IsError 是否返回错误
	IsError() bool
	// Error implement error interface
	Error() string
	// APIRequestID 返回请求ID
	APIRequestID() string
}

// BaseResponse shared api response data fields
type BaseResponse struct {
	// Message 返回信息
	Message string `json:"message"`
	// RequestID 请求的日志id，唯一标识一个请求
	RequestID string `json:"request_id,omitempty"`
	// Code 返回码
	Code int `json:"code"`
}

// IsError implement Response interface
func (r BaseResponse) IsError() bool {
	return r.Code != 0
}

// Error implement Response interface
func (r BaseResponse) Error() string {
	return util.StringsJoin(strconv.Itoa(r.Code), ":", r.Message)
}

// APIRequestID implement Response interface
func (r BaseResponse) APIRequestID() string {
	return r.RequestID
}
