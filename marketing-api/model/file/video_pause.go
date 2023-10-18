package file

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoPauseRequest 按账户暂停素材 API Request
type VideoPauseRequest struct {
	// AdvertiserID 素材归属的广告主
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MaterialIDs 待暂停素材ID
	MaterialIDs []uint64 `json:"material_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r VideoPauseRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// VideoPauseResponse 按账户暂停素材 API Response
type VideoPauseResponse struct {
	model.BaseResponse
	Data *VideoPauseResult `json:"data,omitempty"`
}

type VideoPauseResult struct {
	// MaterialClearRes1ult 暂停结果明细
	MaterialClearResult *MaterialClearResult `json:"material_clear_result,omitempty"`
	// Status 暂停结果, 允许值:
	// SUCCESS 成功
	// PART_SUCCESS 部分成功
	// FAIL 失败
	Status enum.VideoPauseStatus `json:"status,omitempty"`
	// FailMaterialIDs 暂停失败的素材ID(素材ID下仍然有至少一个可投创意)
	FailMaterialIDs []uint64 `json:"fail_material_ids,omitempty"`
}

// MaterialClearResult 暂停结果明细
type MaterialClearResult struct {
	// MaterialID 素材ID
	MaterialID uint64 `json:"material_id,omitempty"`
	// IsCleared 素材是否被暂停（存在内部错误导致素材没有被暂停)
	IsCleared bool `json:"is_cleared,omitempty"`
	// CreativeSuccessList 暂停成功的适用该素材创意ID
	CreativeSuccessList []uint64 `json:"creative_success_list,omitempty"`
	// CreativeFailureList 暂停失败的适用该素材创意ID
	CreativeFailureList []uint64 `json:"creative_failure_list,omitempty"`
	// PromotionSuccessList 暂停成功的适用该素材广告ID
	PromotionSuccessList []uint64 `json:"promotion_success_list,omitempty"`
	// PromotionFailureList 暂停失败的适用该素材广告ID
	PromotionFailureList []uint64 `json:"promotion_failure_list,omitempty"`
}
