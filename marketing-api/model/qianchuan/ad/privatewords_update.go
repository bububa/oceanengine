package ad

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// PrivatewordsUpdateRequest 全量更新否定词 API Request
type PrivatewordsUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	PrivateWords
}

// Encode implement PostRequest interface
func (r PrivatewordsUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// PrivatewordsUpdateResponse 全量更新否定词 API Response
type PrivatewordsUpdateResponse struct {
	model.BaseResponse
	Data *PrivatewordsUpdateResult `json:"data,omitempty"`
}

type PrivatewordsUpdateResult struct {
	// Status 修改状态 可选值:
	// UNKNOWN: 未知结果
	// SUCCESS: 成功
	// FAIL: 失败
	Status string `json:"status,omitempty"`
	// Reason 原因
	Reason string `json:"reason,omitempty"`
}
