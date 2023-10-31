package v3

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// RepeatedDayType 重复周期
type RepeatedDayType string

const (
	PER_MONDAY    RepeatedDayType = "PER_MONDAY"
	PER_TUESDAY   RepeatedDayType = "PER_TUESDAY"
	PER_WEDNESDAY RepeatedDayType = "PER_WEDNESDAY"
	PER_THURSDAY  RepeatedDayType = "PER_THURSDAY"
	PER_FRIDAY    RepeatedDayType = "PER_FRIDAY"
	PER_SATURDAY  RepeatedDayType = "PER_SATURDAY"
	PER_SUNDAY    RepeatedDayType = "PER_SUNDAY"
	EVERY_DAY     RepeatedDayType = "EVERY_DAY"
)

// SetRequest 设置一键起量 API Request
type SetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionID 广告计划id
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// RaiseInfo 起量信息，最多支持6个，如果需要删除请传[]
	RaiseInfo []RaiseInfo `json:"raise_info,omitempty"`
}

// RaiseInfo 起量信息
type RaiseInfo struct {
	// RaiseEndTime 预计结束时间
	RaiseEndTime string `json:"raise_end_time,omitempty"`
	// RaiseBudget 起量预算，单位：元，允许小数点后两位起量预算需大于等于计划出价，小于等于计划预算
	RaiseBudget float64 `json:"raise_budget,omitempty"`
	// IsEffectiveNow 是否立即生效，仅支持广告状态为“投放中”的广告，仅支持1个方案设置“立即生效”，传入True时不支持填写appointed_time
	IsEffectiveNow bool `json:"is_effective_now,omitempty"`
	// AppointedTime 指定投放时间，当is_effective_now为FALSE时填写有效
	AppointedTime *AppointedTime `json:"appointed_time,omitempty"`
}

type AppointedTime struct {
	// RepeatedDay 重复周期，仅生效一次。不传则不重复，传入EVERY_DAY则每天重复
	// 允许值：
	// PER_MONDAY；PER_TUESDAY；PER_WEDNESDAY；PER_THURSDAY；PER_FRIDAY：PER_SATURDAY；PER_SUNDAY；EVERY_DAY
	// EVERY_DAY和其他允许值不可同时传入
	RepeatedDay []RepeatedDayType `json:"repeated_day,omitempty"`
	// RaiseTime 起量时间
	// 重复周期不传时，格式为yyyy-mm-dd HH:MM
	// 传重复周期时，格式为HH:MM
	RaiseTime string `json:"raise_time,omitempty"`
}

// Encode implement PostRequest interface
func (r SetRequest) Encode() []byte {
	if r.RaiseInfo == nil {
		r.RaiseInfo = make([]RaiseInfo, 0)
	}
	return util.JSONMarshal(r)
}
