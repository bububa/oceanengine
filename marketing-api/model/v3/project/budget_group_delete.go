package project

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// BudgetGroupDeleteRequest 批量删除预算组 API Request
type BudgetGroupDeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// 要删除的预算组id列表，请注意：
	// 删除预算组，删除后关联的项目不受影响，继续以项目预算进行推广
	// 预算组删除后不可恢复、不可见
	// 预算组内无项目、项目均为已删除时，才支持删除预算组
	BudgetGroupIDs []uint64 `json:"budget_group_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r BudgetGroupDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// BudgetGroupDeleteResponse 批量删除预算组 API Response
type BudgetGroupDeleteResponse struct {
	model.BaseResponse
	Data *BudgetGroupDeleteResult `json:"data,omitempty"`
}

type BudgetGroupDeleteResult struct {
	// BudgetGroupIDs 删除成功的预算组ID
	BudgetGroupIDs []uint64 `json:"budget_group_ids,omitempty"`
	// Errors 删除失败的预算组
	Errors []BudgetGroupDeleteError `json:"errors,omitempty"`
}

type BudgetGroupDeleteError struct {
	// BudgetGroupID 预算组ID
	BudgetGroupID uint64 `json:"budget_group_id,omitempty"`
	// ErrorMessage 错误信息
	ErrorMessage string `json:"error_message,omitempty"`
}

// Error implement Error interface
func (e BudgetGroupDeleteError) Error() string {
	return util.StringsJoin("广告项目ID:", strconv.FormatUint(e.BudgetGroupID, 10), ", ", e.ErrorMessage)
}
