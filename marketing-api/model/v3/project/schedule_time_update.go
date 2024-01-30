package project

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ScheduleTimeUpdateRequest 批量更新项目投放时段 API request
type ScheduleTimeUpdateRequest struct {
	// AdvertiserID 广告账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Data 批量修改投放时段，限制最多批量修改10条广告
	Data []ScheduleTimeUpdateData `json:"data,omitempty"`
}

// ScheduleTimeUpdateData 批量修改投放时段
type ScheduleTimeUpdateData struct {
	// 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// ScheduleType 投放时间类型，允许值：
	// SCHEDULE_FROM_NOW从今天起长期投放（默认）
	// SCHEDULE_START_END结束日期
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// EndTime 结束的投放时间，当schedule_type 为SCHEDULE_START_END时必填
	EndTime int64 `json:"end_time,omitempty"`
}

// Encode implement PostRequest interface
func (r ScheduleTimeUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
