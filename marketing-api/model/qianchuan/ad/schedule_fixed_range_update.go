package ad

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ScheduleFixedRangeUpdateRequest 更新计划投放时长 API Request
type ScheduleFixedRangeUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdIDs 需要更新的广告计划id，一次只允许传入10个
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// ScheduleFixedRange 固定投放时长
	// 单位为秒，最小值为1800（0.5小时），最大值为48*1800（24小时），值必须为1800倍数，不然会报错
	ScheduleFixedRange int64 `json:"schedule_fixed_range,omitempty"`
}

// Encode implement PostRequest interface
func (r ScheduleFixedRangeUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
