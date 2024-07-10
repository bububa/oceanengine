package promotion

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MaterialStatusUpdateRequest 批量更新广告素材启用状态 API Request
type MaterialStatusUpdateRequest struct {
	// Data 批量更新广告启用状态，包含广告ID和操作状态，list长度限制1～10
	Data []MaterialStatusUpdateItem `json:"data,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionID 广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
}

// MaterialStatusUpdateItem 批量更新广告启用状态，包含广告ID和操作状态
type MaterialStatusUpdateItem struct {
	// OptStatus 允许值：DISABLE 暂停、ENABLE 启用
	OptStatus enum.OptStatus `json:"opt_status,omitempty"`
	// MaterialID 素材ID
	MaterialID uint64 `json:"material_id,omitempty"`
}

// Encode implement PostRequest interface
func (r MaterialStatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// MaterialStatusUpdateResponse 批量更新广告素材启用状态 API Response
type MaterialStatusUpdateResponse struct {
	Data *MaterialStatusUpdateResult `json:"data,omitempty"`
	model.BaseResponse
}

// MaterialStatusUpdateResult
type MaterialStatusUpdateResult struct {
	// MaterialIDs 更新成功的素材集合
	MaterialIDs []uint64 `json:"material_ids,omitempty"`
	// Errors 更新失败的素材列表
	Errors []MaterialStatusUpdateError `json:"errors,omitempty"`
	// PromotionID 广告计划ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
}

// MaterialStatusUpdateError 更新失败的素材
type MaterialStatusUpdateError struct {
	// Message 失败原因
	Message string `json:"message,omitempty"`
	// MaterialID 素材ID
	MaterialID uint64 `json:"material_id,omitempty"`
}

func (e MaterialStatusUpdateError) Error() string {
	return util.StringsJoin("更新素材启用状态失败，素材ID: ", strconv.FormatUint(e.MaterialID, 10))
}
