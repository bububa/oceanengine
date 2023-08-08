package ad

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ScheduleDateUpdateRequest 更新计划投放时间 API Request
type ScheduleDateUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdIDs 需要更新的广告计划id，一次只允许传入10个
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// ScheduleType  投放日期选择方式，允许值：
	// SCHEDULE_FROM_NOW 长期投放
	// SCHEDULE_START_END 具体日期投放
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// EndTime 投放结束时间，形式如：2017-01-01,结束时间不能早于当天。
	// 注意：当video_schedule_type=SCHEDULE_START_END 必填。
	EndTime string `json:"end_time,omitempty"`
}

// Encode implement PostRequest interface
func (r ScheduleDateUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
