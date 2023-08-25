package ad

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ScheduleTimeUpdateRequest 更新计划投放时段 API Request
type ScheduleTimeUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdIDs 需要更新的广告计划id，一次只允许传入10个
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// ScheduleTime 投放时段
	// 默认全时段投放（不限），格式是48*7位字符串，且都是0或1。也就是以半个小时为最小粒度，周一至周日每天分为48个区段，0为不投放，1为投放，不传、全传0、全传1均代表全时段投放。
	// 例如：填写"000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000"，则投放时段为周一到周日的11:30~13:30
	ScheduleTime string `json:"schedule_time,omitempty"`
}

// Encode implement PostRequest interface
func (r ScheduleTimeUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
