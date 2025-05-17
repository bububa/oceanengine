package project

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/local"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateRequest 更新项目 API Request
type UpdateRequest struct {
	// LocalAccountID 本地推广告账户ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// Name 项目名称，长度是1-50个字（两个英文字符占1个字）
	Name string `json:"name,omitempty"`
	// Audience 定向设置，接口暂不支持设置抖音达人定向，该定向默认不限
	Audience *Audience `json:"audience,omitempty"`
	// ScheduleType 投放日期类型设置，允许值：
	// FROM_NOW_ON 从今天起长期投放
	// START_TO_END 设置开始结束时间
	// FIXED_TIME 固定时长
	// 仅营销场景=LIVE直播时，该枚举值有效，否则传入报错
	ScheduleType local.ScheduleType `json:"schedule_type,omitempty"`
	// ScheduleFixedSeconds 直播固定投放时长（单位：秒）
	// 填写说明：
	// 仅当营销场景=LIVE && schedule_type=FIXED_TIME 固定时长时有效且必填，其他情况传入报错
	// 输入值需不小于1800，且为1800的整数倍（即半个小时为最小粒度）
	ScheduleFixedSeconds int64 `json:"schedule_fixed_seconds,omitempty"`
	// EndTime 结束投放时间，精确到天，例如：2017-01-01
	// 当schedule_type=START_TO_END时，有效且必填
	// 结束时间不得小于开始时间
	EndTime string `json:"end_time,omitempty"`
	// ScheduleTime 投放时段，默认全时段投放。格式是48*7位字符串，且都是0或1，也就是以半个小时为最小粒度，周一至周日每天分为48个区段，0为不投放，1为投放，不传、全传0、全传1均代表全时段投放。
	// 当schedule_type=FIXED_TIME固定时长时，该字段无效，传入会报错
	// 例如：000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000，则投放时段为周一到周日的11:30~13:30
	ScheduleTime string `json:"schedule_time,omitempty"`
	// Bid 出价，当出价方式=MANUAL 手动出价时有效
	// 出价单位：分
	// 展示量取值范围为[400,10000]，其他场景取值范围为[1, 1000000]
	// 【展示量】：分/千次曝光
	Bid int64 `json:"bid,omitempty"`
	// Budget 项目预算，单位为：分
	// 填写限制：
	// 当budget_mode=BUDGET_MODE_DAY 日预算 && bid_type=SMART 智能出价时，取值范围为[10000, 999999999]
	// 当budget_mode=BUDGET_MODE_DAY 日预算 && bid_type=MANUAL 手动出价时，取值范围为[30000, 999999999]
	// 当budget_mode=BUDGET_MODE_TOTAL 总预算时，取值范围为[30000, 999999999]
	Budget int64 `json:"budget,omitempty"`
	// IsSetPeakBudget 高峰日预算设置，仅当营销场景为短视频，且推广目的=团购成交或门店引流时，该字段有效且必填，允许值：
	// TURE 开启
	// FALSE 关闭
	// 该字段为FALSE时：高峰日（自然周、节假日）、高峰日预算上调比例 均不可填值
	IsSetPeakBudget string `json:"is_set_peak_budget,omitempty"`
	// PeakWeekDays 高峰日-自然周，允许值：
	// MONDAY 周一
	// TUESDAY 周二
	// WEDNESDAY 周三
	// THURSDAY 周四
	// FRIDAY 周五
	// SATURDAY 周六
	// SUNDAY 周日
	// 注意：peak_holidays和peak_week_days不可同时为空，必须传入至少一个高峰日
	PeakWeekDays []enum.Weekday `json:"peak_week_days,omitempty"`
	// PeakHolidays 高峰日-节假日，允许值：
	// NEW_YEAR 元旦
	// SPRING_FESTIVAL 春节
	// VALENTINES_DAY 情人节
	// LABOUR_DAY 五一
	// DRAGON_BOAT_FESTIVAL 端午
	// CHINESE_VALENTINES_DAY 七夕
	// MOON_FESTIVAL 中秋
	// NATIONAL_DAY 国庆
	// CHRISTMAS 圣诞节
	// SHOPPING_DAY_618 618
	// SHOPPING_DAY_1111 双11
	// 注意：peak_holidays和peak_week_days不可同时为空，必须传入至少一个高峰日
	PeakHolidays []enum.Holiday `json:"peak_holidays,omitempty"`
	// HighBudgetRate 高峰日预算上调比例，单位为百分比，例如：传“40”表示高峰日时预算上调“40%”
	// 当is_set_peak_budget = TURE开启高峰日预算时有效且必填
	// 区间限制： 20～200
	HighBudgetRate int `json:"high_budget_rate,omitempty"`
}

// Encode implements PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
