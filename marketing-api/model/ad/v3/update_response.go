package v3

import "github.com/bububa/oceanengine/marketing-api/model"

// UpdateResponse 更新计划API Response
type UpdateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *UpdateResponseData `json:"data,omitempty"`
}

// UpdateResponseData json返回值
type UpdateResponseData struct {
	// PromotionID 广告计划ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// PromotionIDs 广告计划ID集合
	PromotionIDs []uint64 `json:"promotion_ids,omitempty"`
	// Errors 更新失败的广告计划列表
	Errors []UpdateError `json:"errors,omitempty"`
}

// UpdateError 更新失败的广告计划
type UpdateError struct {
	// PromotionID 广告计划ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// ObjectType 错误对象类型
	// 枚举值： BASIC 广告基本设置、MATERIAL 广告素材组合、BUDGET 广告预算
	ObjectType string `json:"object_type,omitempty"`
	// ErrorCode 错误信息
	ErrorCode int `json:"error_code,omitempty"`
	// ErrorMessage 错误信息
	ErrorMessage string `json:"error_message"`
}
