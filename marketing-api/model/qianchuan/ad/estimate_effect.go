package ad

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// EstimateEffectRequest 获取预估效果接口 API Request
type EstimateEffectRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 抖音id，即商品广告背后关联的抖音号，可通过【查询可推广抖音号列表】接口获取名下可推广抖音号
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// ExternalAction  转化目标
	// 1、直播间成交 AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY
	// 2、支付ROI：external_action=AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY且deep_external_action=AD_CONVERT_TYPE_LIVE_PAY_ROI且deep_bid_type=MIN
	ExternalAction qianchuan.ExternalAction `json:"external_action,omitempty"`
	// DeepExternalAction 深度转化目标，对应千川后台「期待同时优化」
	// 允许值：
	// 1、AD_CONVERT_TYPE_LIVE_PAY_ROI ROI相关
	DeepExternalAction qianchuan.ExternalAction `json:"deep_external_action,omitempty"`
	// DeepBidType 深度出价方式，允许值：MIN等同于PC端，转化目标设置为“支付ROI”
	// 1.仅deep_external_action=AD_CONVERT_TYPE_LIVE_PAY_ROI 时，必填
	//  可选值:
	// MIN
	DeepBidType qianchuan.DeepBidType `json:"deep_bid_type,omitempty"`
	// BidgetMode 预算类型，允许值：
	// BUDGET_MODE_DAY
	// BUDGET_MODE_TOTAL
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 预算，最多支持两位小数。预算范围是300 - 9999999.99
	Budget float64 `json:"budget,omitempty"`
	// LiveScheduleType 直播间投放时段选择方式，允许值：
	// 1、SCHEDULE_FROM_NOW 从今天起长期投放、
	// 2、SCHEDULE_START_END 设置开始和结束日期、
	// 3、SCHEDULE_TIME_FIXEDRANGE 固定时长
	LiveScheduleType enum.LiveScheduleType `json:"live_schedule_type,omitempty"`
	// StartTime 投放起始时间，形式如：2017-01-01，广告投放起始时间不允许修改。
	// live_schedule_type为SCHEDULE_START_END 设置开始和结束日期时需传入。
	// 当live_schedule_type为SCHEDULE_FROM_NOW 、SCHEDULE_TIME_FIXEDRANGE 时不能传入
	StartTime string `json:"start_time,omitempty"`
	// EndTime 投放结束时间，形式如：2017-01-01,结束时间不能比起始时间早。
	// 当live_schedule_type为SCHEDULE_START_END 设置开始和结束日期时需传入。
	// 当live_schedule_type为SCHEDULE_FROM_NOW 、SCHEDULE_TIME_FIXEDRANGE 时不能传入
	EndTime string `json:"end_time,omitempty"`
	// ScheduleTime 投放时段当且 live_schedule_type 为SCHEDULE_START_END和SCHEDULE_FROM_NOW时允许传入，默认全时段投放，格式是48*7位字符串，且都是0或1。也就是以半个小时为最小粒度，周一至周日每天分为48个区段，0为不投放，1为投放，不传、全传0、全传1均代表全时段投放。
	// 例如：填写"000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000"，则投放时段为周一到周日的11:30~13:30
	ScheduleTime string `json:"schedule_time,omitempty"`
	// ScheduleFixedRange 固定投放时长
	// 当 live_schedule_type 为 SCHEDULE_TIME_FIXEDRANGE 固定时长时必填；
	// 当live_schedule_type 为SCHEDULE_FROM_NOW 从今天起长期投放、SCHEDULE_START_END 设置开始和结束日期时不能传入。
	// 单位为秒，最小值为1800（0.5小时），最大值为48*1800（24小时），值必须为1800倍数，不然会报错
	ScheduleFixedRange int64 `json:"schedule_fixed_range,omitempty"`
}

// Encode implement GetRequest interface
func (r EstimateEffectRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("aweme_id", strconv.FormatUint(r.AwemeID, 10))
	values.Set("external_action", string(r.ExternalAction))
	if r.DeepExternalAction != "" {
		values.Set("deep_external_action", string(r.DeepExternalAction))
	}
	if r.DeepBidType != "" {
		values.Set("deep_bid_type", string(r.DeepBidType))
	}
	values.Set("budget_mode", string(r.BudgetMode))
	values.Set("budget", strconv.FormatFloat(r.Budget, 'f', 2, 64))
	values.Set("live_schedule_type", string(r.LiveScheduleType))
	if r.StartTime != "" {
		values.Set("start_time", r.StartTime)
	}
	if r.EndTime != "" {
		values.Set("end_time", r.EndTime)
	}
	if r.ScheduleTime != "" {
		values.Set("schedule_time", r.ScheduleTime)
	}
	if r.ScheduleFixedRange > 0 {
		values.Set("schedule_fixed_range", strconv.FormatInt(r.ScheduleFixedRange, 10))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// EstimateEffectResponse 获取预估效果接口 API Response
type EstimateEffectResponse struct {
	model.BaseResponse
	Data *EstimateEffectResult `json:"data,omitempty"`
}

type EstimateEffectResult struct {
	// EstimateEffectLow 预估结果下限，单位千分之一分
	// 当优化目标为「直播间成交」时，预估转化成本；
	// 当优化目标为「支付ROI」时，预估支付ROI
	EstimateEffectLow float64 `json:"estimate_effect_low,omitempty"`
	// EstimateEffectHigh 预估结果上限，单位千分之一分
	EstimateEffectHigh float64 `json:"estimate_effect_high,omitempty"`
}
