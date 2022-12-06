package adraise

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// EstimateRequest 获取起量预估值 API Request
type EstimateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 广告计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// ModifyValue 预估值，启动一键起量时必填，单位千分之一分，取值大于等于0且小于等于计划预算
	ModifyValue int64 `json:"modify_value,omitempty"`
}

// Encode implement GetRequest interface
func (r EstimateRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("ad_id", strconv.FormatUint(r.AdID, 10))
	if r.ModifyValue > 0 {
		values.Set("modify_value", strconv.FormatInt(r.ModifyValue, 10))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// EstimateResponse 获取起量预估值 API Response
type EstimateResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data struct {
		// EstimateNum 预估展示量
		EstimateNum int64 `json:"estimate_num,omitempty"`
	} `json:"data,omitempty"`
}
