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
	// StartTime 投放开始时间
	// 注意：当schedule_type=SCHEDULE_START_END时必填
	StartTime string `json:"start_time,omitempty"`
	// EndTime 投放结束时间
	// 注意：当schedule_type=SCHEDULE_START_END时必填
	EndTime string `json:"end_time,omitempty"`
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
