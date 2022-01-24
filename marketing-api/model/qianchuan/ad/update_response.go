package ad

import "github.com/bububa/oceanengine/marketing-api/model"

// UpdateResponse 计划更新 API Response
type UpdateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *UpdateResponseData `json:"data,omitempty"`
}

// UpdateResponseData json返回值
type UpdateResponseData struct {
	// AdIDs 更新成功的计划id
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// Errors 更新失败的广告计划列表
	Errors []UpdateError `json:"errors,omitempty"`
}

// UpdateError 更新失败错误
type UpdateError struct {
	// ObjectID 错误对象id
	ObjectID uint64 `json:"object_id,omitempty"`
	// ObjectType 错误对象类型，返回值：AD 计划，CREATIVE 创意
	ObjectType string `json:"object_type,omitempty"`
	// OptType 操作类型，返回值：UPDATE 更新，ADD 新建
	OptType string `json:"opt_type,omitempty"`
	// AdID 更新失败的计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// ErrorCode 错误码
	ErrorCode int `json:"error_code,omitempty"`
	// ErrorMessage 错误信息
	ErrorMessage string `json:"error_message,omitempty"`
}
