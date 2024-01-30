package project

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

type WeekScheduleUpdateRequest struct {
	// AdvertiserID 广告账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Data 批量修改投放时段，限制最多批量修改10条广告
	Data []WeekScheduleUpdateData `json:"data,omitempty"`
}

// WeekScheduleUpdateData 批量修改投放时段
type WeekScheduleUpdateData struct {
	ProjectID uint64 `json:"project_id,omitempty"`
	// ScheduleTime 更新后的投放时段，格式是48*7位字符串，且都是0或1。也就是以半个小时为最小粒度，周一至周日每天分为48个区段，0为不投放，1为投放，不传、全传0、全传1均代表全时段投放。
	// 例如：000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000，则投放时段为周一到周日的11:30~13:30
	// 所选时段必须在广告所属项目的投放时段范围内
	ScheduleTime string `json:"schedule_time,omitempty"`
	// ScheduleScene 生效方式，允许值：
	// NEXT_DAY 次日0点生效
	// REALTIME 立即生效
	//  默认值: REALTIME
	ScheduleScene enum.ScheduleScene `json:"schedule_scene,omitempty"`
}

// Encode implement PostRequest interface
func (r WeekScheduleUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
