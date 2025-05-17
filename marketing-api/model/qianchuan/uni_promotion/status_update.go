package unipromotion

import (
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// StatusUpdateRequest 更改全域推广计划状态 API Request
type StatusUpdateRequest struct {
	// AdvertiserID 千川广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdIDs 全域推广计划id
	// 注意：目前仅支持一个
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// OptStatus 更新推广计划状态，可选值:
	// DISABLE 暂停计划
	// ENABLE 启用计划
	OptStatus qianchuan.AdOptStatus `json:"opt_status,omitempty"`
}

// Encode implements PostRequest interface
func (r StatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// StatusUpdateResponse 编辑全域推广计划 API Response
type StatusUpdateResponse struct {
	model.BaseResponse
	Data struct {
		// Results 更新结果
		Results []StatusUpdateResult `json:"results,omitempty"`
	} `json:"data,omitempty"`
}

type StatusUpdateResult struct {
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// Flag 是否更新成功
	Flag bool `json:"flag,omitempty"`
	// Error 错误信息
	Error *Error `json:"error,omitempty"`
}
