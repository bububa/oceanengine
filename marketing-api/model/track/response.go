package track

import (
	"fmt"
)

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
