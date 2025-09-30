package unipromotion

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdScheduleDateUpdateRequest 更新全域推广计划投放时间 API Request
type AdScheduleDateUpdateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UpdateScheduleInfos 更新投放时间信息
	// 注意：单次最多支持10个
	UpdateScheduleInfos []UpdateScheduleInfo `json:"update_schedule_infos,omitempty"`
}

// UpdateScheduleInfo 更新投放时间信息
type UpdateScheduleInfo struct {
	// AdID 商品全域计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// ScheduleType 全域投放时间选择方式，可选值:
	// SCHEDULE_FROM_NOW 从今天起长期投放
	// SCHEDULE_START_END 设置开始和结束日期
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// EndTime 投放结束时间，形式如：2017-01-01,结束时间不能比起始时间/当前时间早
	// 注意：当schedule_type=SCHEDULE_START_END时必填；当schedule_type=SCHEDULE_FROM_NOW时无需入参
	EndTime string `json:"end_time,omitempty"`
}

// Encode implements PostRequest inteface
func (r AdScheduleDateUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
