package adraise

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SetRequest 设置一键起量 API Request
type SetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 广告计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// OptType 操作类型
	OptType enum.AdRaiseOptType `json:"opt_type,omitempty"`
	// ModifyValue 预估值，启动一键起量时必填，单位千分之一分，取值大于等于0且小于等于计划预算
	ModifyValue int64 `json:"modify_value,omitempty"`
}

// Encode implement PostRequest interface
func (r SetRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// SetResponse 设置一键起量 API Response
type SetResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data struct {
		// Status 一键起量状态
		Status enum.AdRaiseStatus `json:"status,omitempty"`
	} `json:"data,omitempty"`
}
