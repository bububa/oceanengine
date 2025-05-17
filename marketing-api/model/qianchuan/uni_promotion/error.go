package unipromotion

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/util"
)

// Error 错误信息
type Error struct {
	// ObjectID 错误对象id
	ObjectID uint64 `json:"object_id,omitempty"`
	// ObjectType 错误对象类型
	ObjectType int `json:"object_type,omitempty"`
	// OptType 操作类型
	OptType int `json:"opt_type,omitempty"`
	// ErrorCode 错误码
	ErrorCode int `json:"error_code,omitempty"`
	// ErrorMessage 错误信息
	ErrorMessage string `json:"error_message,omitempty"`
	// Extra 额外信息
	Extra map[string]interface{} `json:"extra,omitempty"`
}

// Error implements error interface
func (e Error) Error() string {
	ret := util.StringsJoin("code: ", strconv.Itoa(e.ErrorCode), ", msg:", e.ErrorMessage)
	if e.ObjectID > 0 {
		ret = util.StringsJoin("object: ", strconv.FormatUint(e.ObjectID, 10), ", ", ret)
	}
	return ret
}
