package unipromotion

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateRequest 编辑全域推广计划 API Request
type UpdateRequest struct {
	// DeliverySetting 投放设置
	DeliverySetting *DeliverySetting `json:"delivery_setting,omitempty"`
	// CreativeSetting 创意设置
	CreativeSetting *CreativeSetting `json:"creative_setting,omitempty"`
	// ProgrammaticCreativeMediaList 程序化创意信息
	ProgrammaticCreativeMediaList []ProgrammaticCreativeMedia `json:"programmatic_creative_media_list,omitempty"`
	// AdvertiserID 千川广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 全域推广计划id
	AdID uint64 `json:"ad_id,omitempty"`
}

// Encode implements PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 编辑全域推广计划 API Response
type UpdateResponse struct {
	Data *UpdateResult `json:"data,omitempty"`
	model.BaseResponse
}

type UpdateResult struct {
	// ErrorList 错误信息
	ErrorList []Error `json:"error_list,omitempty"`
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
}
