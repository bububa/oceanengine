package ad

import "github.com/bububa/oceanengine/marketing-api/model"

// UpdateResponse 更新计划API Response
type UpdateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *UpdateResponseData `json:"data,omitempty"`
}

// UpdateResponseData json返回值
type UpdateResponseData struct {
	// AdIDs 广告计划ID集合
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// Errors 更新失败的广告计划列表
	Errors []UpdateError `json:"errors,omitempty"`
}

// UpdateError 更新失败的广告计划
type UpdateError struct {
	// AdID 广告计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// ErrorMessage 错误信息
	ErrorMessage string `json:"error_message"`
}
