package ad

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
)

// DeliverySetting 投放设置
type DeliverySetting struct {
	// SmartBidType 投放场景（出价方式）
	SmartBidType enum.SmartBidType `json:"smart_bid_type,omitempty"`
	// FlowControlMode 投放速度
	FlowControlMode enum.FlowControlMode `json:"flow_control_mode,omitempty"`
	// ExternalAction 转化目标
	ExternalAction qianchuan.ExternalAction `json:"external_action,omitempty"`
	// DeepExternalAction 深度转化目标
	DeepExternalAction qianchuan.ExternalAction `json:"deep_external_action,omitempty"`
	// DeepBidType 深度出价方式，允许值：MIN
	DeepBidType qianchuan.DeepBidType `json:"deep_bid_type,omitempty"`
	// RoiGoal 支付ROI目标，最多支持两位小数，0.01～100
	// 当external_action=AD_CONVERT_TYPE_SHOPPING且deep_external_action=AD_CONVERT_TYPE_LIVE_PAY_ROI且deep_bid_type=MIN时，必填
	// 注意：
	// 1、按展示付费(oCPM)，根据【保障规则】提供保障福利，请谨慎修改支付ROI目标和定向，以免失去保障资格。
	RoiGoal float64 `json:"roi_goal,omitempty"`
	// Budget 预算
	Budget float64 `json:"budget,omitempty"`
	// ReviveBudget 复活预算
	ReviveBudget float64 `json:"revive_budget,omitempty"`
	// BudgetMode 预算类型
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// CpaBid 转化出价
	CpaBid float64 `json:"cpa_bid,omitempty"`
	// VideoScheduleType 短视频投放日期选择方式
	VideoScheduleType enum.VideoScheduleType `json:"video_schedule_type,omitempty"`
	// LiveScheduleType 直播间投放时段选择方式
	LiveScheduleType enum.LiveScheduleType `json:"live_schedule_type,omitempty"`
	// StartTime 投放开始时间
	StartTime string `json:"start_time,omitempty"`
	// EndTime 投放结束时间
	EndTime string `json:"end_time,omitempty"`
	// ScheduleTime 投放时段，当 video_schedule_type 和 live_schedule_type为SCHEDULE_START_END和SCHEDULE_FROM_NOW时有值，格式是48*7位字符串，且都是0或1。也就是以半个小时为最小粒度，周一至周日每天分为48个区段，0为不投放，1为投放，不传、全传0、全传1均代表全时段投放。
	ScheduleTime string `json:"schedule_time,omitempty"`
	// ScheduleFixedRange 固定投放时长，当 live_schedule_type 为时有值；单位为秒，最小值为1800（0.5小时），最大值为48*1800（24小时）SCHEDULE_TIME_FIXEDRANGE
	ScheduleFixedRange int64 `json:"schedule_fixed_range,omitempty"`
	// EnableAutoPause 是否启用超成本自动暂停，允许值：
	// 0 关闭 1 开启
	EnableAutoPause int `json:"enable_auto_pause,omitempty"`
	// AutoManageStrategyCmd 托管策略，允许值：
	// 0 优先跑量 1 优先成本
	AutoManageStrategyCmd int `json:"auto_manage_strategy_cmd,omitempty"`
	// EnableFlowMaterial 是否优质素材自动同步投放
	// 0 关闭 1 开启
	EnableFlowMaterial int `json:"enable_flow_material,omitempty"`
	// ProductNewOpen 是否开启新品加速，默认关闭。
	// 开启：true
	// 关闭：false（默认）
	// 注意：需要调用【商家获取可投商品列表】/【达人获取可投商品列表】获取当前商品是否支持开启新品加速。在开启商品加速后，roi_goal会受到限制。
	ProductNewOpen bool `json:"product_new_open,omitempty"`
	// QcpxMode 是否开启智能优惠券，允许值
	// QCPX_MODE_ON 启用
	// QCPX_MODE_OFF 不启用
	QcpxMode qianchuan.QcpxMode `json:"qcpx_mode,omitempty"`
	// AllowQcpx 是否支持智能优惠券
	// true 支持
	// false 不支持
	AllowQcpx bool `json:"allow_qcpx,omitempty"`
}
