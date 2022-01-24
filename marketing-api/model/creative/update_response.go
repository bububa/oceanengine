package creative

import "github.com/bububa/oceanengine/marketing-api/model"

type UpdateResponse struct {
	model.BaseResponse
	Data *UpdateResponseData `json:"data,omitempty"`
}

type UpdateResponseData struct {
	// Success 更新状态成功的创意ID列表
	Success []uint64 `json:"success,omitempty"`
	// Errors 更新失败的创意列表
	Errors []UpdateError `json:"errors,omitempty"`
}

type UpdateError struct {
	// CreativeID 更新失败的创意ID
	CreativeID uint64 `json:"creative_id,omitempty"`
	// ErrorMessage 更新失败的原因
	ErrorMessage string `json:"error_message,omitempty"`
}
