package project

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// StatusUpdateRequest 批量更新项目状态 API Request
type StatusUpdateRequest struct {
	// LocalAccountID 本地推广告账户ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// Data 批量更新项目状态，包含项目ID和目标操作，list长度限制1～50
	Data []StatusUpdateItem `json:"data,omitempty"`
}

// StatusUpdateItem 批量更新项目
type StatusUpdateItem struct {
	// ProjectID 项目id
	ProjectID uint64 `json:"project_id,omitempty"`
	// OptStatus 目标操作
	// 目标操作，可选值:
	// ENABLE 启用项目
	// PAUSED 暂停项目
	// 对于删除的广告项目不可进行任何操作，否则会报错
	OptStatus string `json:"opt_status,omitempty"`
}

// Encode implements PostRequest interface
func (r StatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// StatusUpdateResponse 批量更新项目状态 API Response
type StatusUpdateResponse struct {
	model.BaseResponse
	Data *UpdateResult `json:"data,omitempty"`
}

type UpdateResult struct {
	// ProjectIDs 更新成功的项目ID
	ProjectIDs []uint64 `json:"project_ids,omitempty"`
	// Errors 更新失败的项目列表
	Errors []UpdateError `json:"errors,omitempty"`
}

// UpdateError 更新失败的项目列表
type UpdateError struct {
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// ErrorMessage 失败信息
	ErrorMessage string `json:"error_message,omitempty"`
}

func (e UpdateError) Error() string {
	return util.StringsJoin("项目ID:", strconv.FormatUint(e.ProjectID, 10), ", ", e.ErrorMessage)
}
