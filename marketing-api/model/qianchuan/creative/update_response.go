package creative

import "github.com/bububa/oceanengine/marketing-api/model"

// UpdateResponse 计划更新 API Response
type UpdateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *UpdateResponseData `json:"data,omitempty"`
}

// UpdateResponseData json返回值
type UpdateResponseData struct {
	// CreativeIDs 更新成功的创意id
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	// Errors 更新失败的广告计划列表
	Errors []UpdateError `json:"errors,omitempty"`
}

// UpdateError 更新失败错误
type UpdateError struct {
	// CreativeID 更新失败的创意id
	CreativeID uint64 `json:"creative_id,omitempty"`
	// ErrorMessage 错误信息
	ErrorMessage string `json:"error_message,omitempty"`
}
