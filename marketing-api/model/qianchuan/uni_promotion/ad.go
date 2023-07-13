package unipromotion

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/live"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// Ad 全域推广
type Ad struct {
	// AdInfo 广告信息
	AdInfo *AdInfo `json:"ad_info,omitempty"`
	// RoomInfo 主播信息
	RoomInfo []live.Room `json:"room_info,omitempty"`
	// StatsInfo 消耗指标
	StatsInfo *report.UniPromotionStats `json:"stats_info,omitempty"`
}

// AdInfo 广告信息
type AdInfo struct {
	// ID 推广id
	ID uint64 `json:"id,omitempty"`
	// StartTime 当前周期开始时间，用来返回周期内数据
	StartTime string `json:"start_time,omitempty"`
	// EndTime 当前周期结束时间
	EndTime string `json:"end_time,omitempty"`
	// ModifyTime 修改时间
	ModifyTime string `json:"modify_time,omitempty"`
	// CreateTime 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// MarketingGoal  营销目标
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// Roi2Goal 支付ROI目标，最多支持两位小数
	Roi2Goal float64 `json:"roi2_goal,omitempty"`
	// BudgetMode  预算类型
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 预算
	Budget float64 `json:"budget,omitempty"`
	// Status 投放状态
	Status qianchuan.AdStatus `json:"status,omitempty"`
	// OptStatus 操作状态，详见【附录-枚举值】
	OptStatus qianchuan.AdOptStatus `json:"opt_status,omitempty"`
	// DeliverySecond 投放时长
	DeliverySecond int64 `json:"delivery_second,omitempty"`
}
