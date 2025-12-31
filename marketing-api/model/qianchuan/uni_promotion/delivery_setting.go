package unipromotion

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
)

// DeliverySetting 投放设置
type DeliverySetting struct {
	// SmartBidType 投放场景（出价方式），可选值:
	// SMART_BID_CUSTOM 控成本投放
	SmartBidType enum.SmartBidType `json:"smart_bid_type,omitempty"`
	// Roi2Goal 支付ROI目标，单位元，最多支持两位小数
	Roi2Goal float64 `json:"roi2_goal,omitempty"`
	// QcpxMode 是否开启智能优惠券，可选值:
	// QCPX_MODE_OFF 关闭
	// QCPX_MODE_ON 开启
	QcpxModel qianchuan.QcpxMode `json:"qcpx_mode,omitempty"`
	// Budget 预算，单位元，最多支持两位小数
	Budget float64 `json:"budget,omitempty"`
	// LiveScheduleType 投放时间选择方式，可选值:
	// SCHEDULE_FROM_NOW从今天起长期投放
	// SCHEDULE_START_END设置开始和结束日期
	LiveScheduleType enum.LiveScheduleType `json:"live_schedule_type,omitempty"`
	// VideoScheduleType 商品全域投放时间选择方式，可选值:
	// SCHEDULE_FROM_NOW从今天起长期投放
	// SCHEDULE_START_END设置开始和结束日期
	VideoScheduleType enum.VideoScheduleType `json:"video_schedule_type,omitempty"`
	// StartTime 投放开始时间
	// 注意：当schedule_type=SCHEDULE_START_END时必填
	StartTime string `json:"start_time,omitempty"`
	// EndTime 投放结束时间
	// 注意：当schedule_type=SCHEDULE_START_END时必填
	EndTime string `json:"end_time,omitempty"`
	// DailyDeliveryTime 每日投放时长，0.5h～24.0h，步进0.5h
	// 注意：仅当marketing_goal=LIVE_PROM_GOODS&smart_bid_type=SMART_BID_CONSERVATIVE时支持，其余场景报错
	DailyDeliveryTime float64 `json:"daily_delivery_time,omitempty"`
	// MinEstimateConvert 预估最小转化
	// 注意：
	// 可通过【获取全域建议预算】接口获取预估值， 如果需要生效成本保障，需要传入正确的预估建议。
	// 如果传入的预算为建议预算，该字段必填
	MinEstimateConvert int64 `json:"min_estimate_convert,omitempty"`
	// EstimateConvert 预估转化
	// 注意：
	// 可通过【获取全域建议预算】接口获取预估值， 如果需要生效成本保障，需要传入正确的预估建议。
	// 如果传入的预算为建议预算，该字段必填
	EstimateConvert int64 `json:"estimate_convert,omitempty"`
	// EstimateROIGoal 预估roi
	// 注意：
	// 可通过【获取全域建议预算】接口获取预估值， 如果需要生效成本保障，需要传入正确的预估建议。
	// 如果传入的预算为建议预算，该字段必填
	EstimateROIGoal float64 `json:"estimate_roi_goal,omitempty"`
	// MinEstimateROIGoal 预估最小roi
	// 注意：
	// 可通过【获取全域建议预算】接口获取预估值， 如果需要生效成本保障，需要传入正确的预估建议。
	// 如果传入的预算为建议预算，该字段必填
	MinEstimateROIGoal float64 `json:"min_estimate_roi_goal,omitempty"`
	// ExternalAction 转化目标
	// AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY 直播间成交
	ExternalAction qianchuan.ExternalAction `json:"external_action,omitempty"`
	// DeepExternalAction 深度转化目标
	// AD_CONVERT_TYPE_LIVE_PAY_ROI 支付ROI
	DeepExternalAction qianchuan.ExternalAction `json:"deep_external_action,omitempty"`
	// DeepBidType 深度出价方式
	DeepBidType qianchuan.DeepBidType `json:"deep_bid_type,omitempty"`
	// PricingType 出价方式，可选值:
	// OCPM OCPM 按照展示进行计费
	PricingType enum.PricingType `json:"pricing_type,omitempty"`
	// BudgetMode  预算类型，可选值:
	// BUDGET_MODE_DAY 日预算
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
}
