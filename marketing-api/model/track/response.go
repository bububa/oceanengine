package track

import (
	"fmt"
)

// Response 线索-API上报数据 API Response
type Response struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Ret  int    `json:"ret,omitempty"`
}

func (r Response) IsError() bool {
	return r.Code != 0
}

func (r Response) Error() string {
	return fmt.Sprintf("%d:%s", r.Code, r.Msg)
}
