package project

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/local"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建项目 API Request
type CreateRequest struct {
	// LocalAccountID 本地推广告账户ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// Name 项目名称，长度是1-50个字（两个英文字符占1个字）
	Name string `json:"name,omitempty"`
	// MarketingGoal 营销场景，允许值：
	// LIVE 直播
	// VIDEO_IMAGE 短视频/图文
	MarketingGoal local.MarketingGoal `json:"marketing_goal,omitempty"`
	// LocalDeliveryScene 推广目的，允许值：
	// CONTENT_HEAT 内容加热
	// POI_RECOMMEND 门店引流
	// 当marketing_goal =LIVE 直播时，不支持传入POI_RECOMMEND 门店引流
	// PRODUCT_PAY 团购成交
	LocalDeliveryScene local.LocalDeliveryScene `json:"local_delivery_scene,omitempty"`
	// AdType 广告类型，允许值：
	// GENERAL 通投广告
	// SEARCHING 搜索广告
	// 当marketing_goal=VIDEO_IMAGE 短视频/图文 && local_delivery_scene=CONTENT_HEAT 内容加热 时，不支持传入SEARCHING
	AdType local.AdType `json:"ad_type,omitempty"`
	// DeliveryGoal 投放内容，允许值：
	// POI 门店
	// PRODUCT 商品
	// 填写说明：
	// 仅marketing_goal=VIDEO_AND_IMAGE时，该字段生效且必填；当marketing_goal=LIVE时，该字段无效
	// 当marketing_goal=VIDEO_AND_IMAGE且
	// local_delivery_scene=POI_RECOMMEND 门店引流时，仅支持POI
	DeliveryGoal local.DeliveryGoal `json:"delivery_goal,omitempty"`
	// DeliveryPoiMode 是否投放全部门店，允许值：
	// ALL 投放全部门店
	// PART 投放指定门店，需同时传入promotion_poi_ids
	// 填写说明：
	// 当 marketing_goal=VIDEO_AND_IMAGE且delivery_goal=POI时，有效且必传
	DeliveryPoiMode local.DeliveryPoiMode `json:"delivery_poi_mode,omitempty"`
	// PromotionPoiIDs 推广门店ID列表
	// 填写说明：
	// 当delivery_goal=POI且 delivery_poi_mode=CUSTOM 投放指定门店时，有效且必填
	PromotionPoiIDs []uint64 `json:"promotion_poi_ids,omitempty"`
	// ProductID 推广商品ID，可通过【获取可投商品列表】接口查询获取
	// 填写说明：
	// 当marketing_goal=VIDEO_AND_IMAGE 且 delivery_goal=PRODUCT 时有效且必填
	ProductID uint64 `json:"product_id,omitempty"`
	// AwemeID 用于推广直播间的抖音号，可通过【获取本地推创编可用抖音号】接口获取
	// 填写说明：
	// 当marketing_goal=LIVE时有效且必填
	AwemeID string `json:"aweme_id,omitempty"`
	// ExternalAction 优化目标
	// 填写说明：
	// 当marketing_goal=VIDEO_AND_IMAGE短视频时：
	// local_delivery_scene=PRODUCT_PURCHASE团购成交或POI_ATTRACTION门店引流时，不支持传入优化目标
	// local_delivery_scene=CONTENT_HEAT内容加热时，允许值：
	// NATIVE_ACTION 用户互动
	// FOLLOW_ACTION 粉丝增长
	// SHOW展示量
	// 当marketing_goal=LIVE 直播：
	// local_delivery_scene=PRODUCT_PURCHASE团购成交时：
	// 当ad_type= GENERAL 通投广告时，允许值：
	// LIVE_OTO_GROUP_BUYING 直播间团购购买
	// LIVE_OTO_CLICK 商品点击
	// 当ad_type= SEARCHING搜索广告时，允许值：
	// LIVE_OTO_GROUP_BUYING 直播间团购购买
	// local_delivery_scene=CONTENT_HEAT 内容加热时：
	// 当ad_type= GENERAL 通投广告时，允许值：
	// LIVE_ENGAGE 直播加热
	// FOLLOWER_COUNT 粉丝增长
	// SHOW 展示
	// 当ad_type= SEARCHING 搜索广告时，允许值：
	// LIVE_ENTER_ACTION直播间观看
	// LIVE_STAY_TIME 直播间停留
	ExternalAction local.ExternalAction `json:"external_action,omitempty"`
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
	// StartTime 开始投放时间，精确到天，例如：2017-01-01
	// 当schedule_type=START_TO_END时，有效且必填
	// 广告投放起始时间不允许修改
	// 开始时间不得大于结束时间，且开始时间内不得小于今天
	StartTime string `json:"start_time,omitempty"`
	// EndTime 结束投放时间，精确到天，例如：2017-01-01
	// 当schedule_type=START_TO_END时，有效且必填
	// 结束时间不得小于开始时间
	EndTime string `json:"end_time,omitempty"`
	// ScheduleTime 投放时段，默认全时段投放。格式是48*7位字符串，且都是0或1，也就是以半个小时为最小粒度，周一至周日每天分为48个区段，0为不投放，1为投放，不传、全传0、全传1均代表全时段投放。
	// 当schedule_type=FIXED_TIME固定时长时，该字段无效，传入会报错
	// 例如：000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000，则投放时段为周一到周日的11:30~13:30
	ScheduleTime string `json:"schedule_time,omitempty"`
	// BidType 出价方式，允许值：
	// MANUAL 手动出价
	// SMART 智能出价
	// 填写说明：
	// 优化目标= SHOW 展示量 时，仅支持MANUAL 手动出价
	// 营销目标=LIVE && 推广目的= CONTENT_HEAT内容加热 && 优化目标=LIVE_ENGAGEMENT时，仅支持SMART 智能出价
	BidType local.BidType `json:"bid_type,omitempty"`
	// Bid 出价，当出价方式=MANUAL 手动出价时有效
	// 出价单位：分
	// 展示量取值范围为[400,10000]，其他场景取值范围为[1, 1000000]
	// 【展示量】：分/千次曝光
	Bid int64 `json:"bid,omitempty"`
	// BudgetMode 预算模式设置，允许值：
	// BUDGET_MODE_DAY 日预算
	// BUDGET_MODE_TOTAL 总预算
	// 填写说明：
	// 当sceduale_type=FIXED_TIME固定时长时，仅支持BUDGET_MODE_TOTAL总预算
	// 当sceduale_type=FROM_NOW_ON或START_TO_END时，仅支持BUDGET_MODE_DAY日预算
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
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
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建项目 API Response
type CreateResponse struct {
	model.BaseResponse
	Data struct {
		// ProjectID 项目ID
		ProjectID uint64 `json:"project_id,omitempty"`
	} `json:"data,omitempty"`
}
