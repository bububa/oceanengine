package creative

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateResponse 更新创意 API Response
type UpdateResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *UpdateResponseData `json:"data,omitempty"`
}

// UpdateResponseData json 返回值
type UpdateResponseData struct {
	// Success 更新状态成功的创意ID列表
	Success []uint64 `json:"success,omitempty"`
	// Errors 更新失败的创意列表
	Errors []UpdateError `json:"errors,omitempty"`
}

// UpdateError 更新失败信息
type UpdateError struct {
	// CreativeID 更新失败的创意ID
	CreativeID uint64 `json:"creative_id,omitempty"`
	// ErrorMessage 更新失败的原因
	ErrorMessage string `json:"error_message,omitempty"`
}

// Error implement error interface
func (r UpdateError) Error() string {
	return util.StringsJoin("广告创意ID:", strconv.FormatUint(r.CreativeID, 10), ", 错误信息", r.ErrorMessage)
}
