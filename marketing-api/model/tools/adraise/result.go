package adraise

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ResultRequest 获取一键起量的后验数据 API Request
type ResultRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 广告计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// OptType 操作类型
	OptType enum.AdRaiseOptType `json:"opt_type,omitempty"`
}

// Encode implement GetRequest interface
func (r ResultRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("ad_id", strconv.FormatUint(r.AdID, 10))
	values.Set("opt_type", string(r.OptType))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ResultResponse 获取一键起量的后验数据 API Response
type ResultResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *Result `json:"data,omitempty"`
}

// Result 一键起量的后验数据
type Result struct {
	// StartTime 一键起量开始时间
	StartTime string `json:"start_time,omitempty"`
	// EndTime 一键起量结束时间
	EndTime string `json:"end_time,omitempty"`
	// Cost 一键起量阶段产生消耗
	Cost int64 `json:"cost,omitempty"`
	// Show 一键起量阶段产生展示
	Show int64 `json:"show,omitempty"`
	// Click 一键起量阶段产生点击数
	Click int64 `json:"click,omitempty"`
	// Convert 一键起量阶段产生转换数
	Convert int64 `json:"convert,omitempty"`
	// Ctr CTR 一键起量期间点击率
	Ctr float64 `json:"ctr,omitempty"`
	// Cvr CVR 一键起量期间转化率
	Cvr float64 `json:"cvr,omitempty"`
}
