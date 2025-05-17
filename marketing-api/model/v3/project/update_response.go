package project

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateResponse 更新计划API Response
type UpdateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *UpdateResponseData `json:"data,omitempty"`
}

// UpdateResponseData json返回值
type UpdateResponseData struct {
	// ProjectID 广告项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// PromotionIDs 广告项目ID集合
	ProjectIDs []uint64 `json:"project_ids,omitempty"`
	// ErrorList 更新失败的广告计划列表
	ErrorList []UpdateError `json:"error_list,omitempty"`
	// Errors 更新失败的广告计划列表
	Errors []UpdateError `json:"errors,omitempty"`
	// ErrorKeywordList
	ErrorKeywordList []ErrorKeyword `json:"error_keywords_list,omitempty"`
}

// UpdateError 更新失败的广告项目
type UpdateError struct {
	// ProjectID 广告项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// ObjectType 错误对象类型
	// 枚举值： BASIC 广告基本设置、MATERIAL 广告素材组合、BUDGET 广告预算
	ObjectType string `json:"object_type,omitempty"`
	// ErrorCode 错误信息
	ErrorCode int `json:"error_code,omitempty"`
	// ErrorMessage 错误信息
	ErrorMessage string `json:"error_message"`
}

type ErrorKeyword struct {
	// ErrorKeyword 失败的关键词
	ErrorKeyword string `json:"error_keyword,omitempty"`
	// ErrorMessage 失败原因
	ErrorMessage string `json:"error_message,omitempty"`
}

func (e ErrorKeyword) Error() string {
	return util.StringsJoin("Keyword:", e.ErrorKeyword, ", Message:", e.ErrorMessage)
}

// Error implement error interface
func (r UpdateError) Error() string {
	var objectType string
	switch r.ObjectType {
	case "BASIC":
		objectType = "广告基本设置"
	case "MATERIAL":
		objectType = "广告素材组合"
	case "BUDGET":
		objectType = "广告预算"
	}
	ret := util.StringsJoin("code:", strconv.Itoa(r.ErrorCode), ", msg:", r.ErrorMessage)
	if objectType != "" {
		ret = util.StringsJoin("更新 ", objectType, "失败, ", ret)
	} else if r.ProjectID > 0 {
		ret = util.StringsJoin("广告项目ID:", strconv.FormatUint(r.ProjectID, 10), ", ", ret)
	}
	return ret
}
