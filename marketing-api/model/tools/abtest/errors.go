package abtest

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/util"
)

type Error struct {
	// TestID 关停失败的实验列表
	TestID uint64 `json:"test_ids,omitempty"`
	// Message 失败原因
	Message string `json:"message,omitempty"`
}

func (e Error) Error() string {
	return util.StringsJoin(e.Message, "(", strconv.FormatUint(e.TestID, 10), ")")
}
