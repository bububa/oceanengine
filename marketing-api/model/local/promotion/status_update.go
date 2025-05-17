package promotion

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// StatusUpdateRequest 批量更新广告状态 API Request
type StatusUpdateRequest struct {
	// LocalAccountID 本地推广告账户ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// Data 批量更新广告状态，包含广告ID和目标操作，list长度限制1～50
	Data []StatusUpdateItem `json:"data,omitempty"`
}

// StatusUpdateItem 批量更新广告
type StatusUpdateItem struct {
	// PromotionID 广告id
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// OptStatus 目标操作
	// 目标操作，可选值:
	// ENABLE 启用广告
	// PAUSED 暂停广告
	// 对于删除的广告广告不可进行任何操作，否则会报错
	OptStatus string `json:"opt_status,omitempty"`
}

// Encode implements PostRequest interface
func (r StatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// StatusUpdateResponse 批量更新广告状态 API Response
type StatusUpdateResponse struct {
	model.BaseResponse
	Data *UpdateResult `json:"data,omitempty"`
}

type UpdateResult struct {
	// PromotionIDs 更新成功的广告ID
	PromotionIDs []uint64 `json:"promotion_ids,omitempty"`
	// Errors 更新失败的广告列表
	Errors []UpdateError `json:"errors,omitempty"`
}

// UpdateError 更新失败的广告列表
type UpdateError struct {
	// PromotionID 广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// ErrorMessage 失败信息
	ErrorMessage string `json:"error_message,omitempty"`
}

func (e UpdateError) Error() string {
	return util.StringsJoin("广告ID:", strconv.FormatUint(e.PromotionID, 10), ", ", e.ErrorMessage)
}
