package project

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/local"
	"github.com/bububa/oceanengine/marketing-api/model/local/product"
)

// Project 项目
type Project struct {
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// ProjectStatusFirst 项目一级状态
	ProjectStatusFirst local.ProjectStatus `json:"project_status_first,omitempty"`
	// ProjectStatusSecond 项目二级状态
	ProjectStatusSecond []local.ProjectStatus `json:"project_status_second,omitempty"`
	// LocalAccountID 本地推广告账户ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// AdType 广告类型
	AdType local.AdType `json:"ad_type,omitempty"`
	// MarketingGoal 营销场景
	MarketingGoal local.MarketingGoal `json:"marketing_goal,omitempty"`
	// Name 项目名称
	Name string `json:"name,omitempty"`
	// ProjectCreateTime 项目创建时间，格式：yyyy-MM-dd HH:mm:ss
	ProjectCreateTime string `json:"project_create_time,omitempty"`
	// ProjectModifyTime 项目更新时间，格式：yyyy-MM-dd HH:mm:ss
	ProjectModifyTime string `json:"project_modify_time,omitempty"`
	// ExternalAction 优化目标，枚举值：
	// FOLLOW_ACTION 粉丝增长
	// LIVE_ENGAGEMENT 直播加热
	// LIVE_ENTER_ACTION 直播间观看
	// LIVE_OTO_CLICK 直播间商品点击
	// LIVE_OTO_GROUP_BUYING 直播间团购购买
	// LIVE_STAY_TIME 直播间停留NATIVE_ACTION 用户互动
	// SHOW 展示量
	// OTO_PAY 团购购买
	// POI_RECOMMEND 门店引流
	ExternalAction local.ExternalAction `json:"external_action,omitempty"`
	// DeliveryGoal 投放内容，枚举值：
	// POI 门店
	// PRODUCT 商品
	DeliveryGoal local.DeliveryGoal `json:"delivery_goal,omitempty"`
	// ProjectBudgetMode 项目预算类型，枚举值：
	// BUDGET_MODE_DAY 日预算
	// BUDGET_MODE_TOTAL 总预算
	ProjectBudgetMode enum.BudgetMode `json:"project_budget_mode,omitempty"`
	// ProjectBudget 项目预算，单位为：分
	ProjectBudget int64 `json:"project_budget,omitempty"`
	// BidType 出价方式，枚举值：
	// MANUAL 手动出价
	// SMART 智能出价
	BidType local.BidType `json:"bid_type,omitempty"`
	// ProjectBid 项目出价，单位为：分
	ProjectBid int64 `json:"project_bid,omitempty"`
	// StartTime 投放开始时间，如：2017-01-01 精确到天
	StartTime string `json:"start_time,omitempty"`
	// EndTime 投放结束时间，如：2017-01-01 精确到天
	EndTime string `json:"end_time,omitempty"`
	// PoiInfo 项目推广门店信息
	PoiInfo *PoiInfo `json:"poi_info,omitempty"`
	// ProductInfo 商品信息
	ProductInfo *product.Product `json:"product_info,omitempty"`
	// DeliveryPoiMode 本地投放模式
	DeliveryPoiMode local.DeliveryPoiMode `json:"delivery_poi_mode,omitempty"`
}

// ProjectDetail 项目详情
type ProjectDetail struct {
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// ProjectStatusFirst 项目一级状态
	ProjectStatusFirst local.ProjectStatus `json:"project_status_first,omitempty"`
	// ProjectStatusSecond 项目二级状态
	ProjectStatusSecond []local.ProjectStatus `json:"project_status_second,omitempty"`
	// LocalAccountID 本地推广告账户ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// AwemeID 抖音号ID
	AwemeID string `json:"aweme_id,omitempty"`
	// AdType 广告类型
	AdType local.AdType `json:"ad_type,omitempty"`
	// LocalDeliveryScene 本地投放场景
	LocalDeliveryScene local.LocalDeliveryScene `json:"local_delivery_scene,omitempty"`
	// MarketingGoal 营销场景
	MarketingGoal local.MarketingGoal `json:"marketing_goal,omitempty"`
	// Name 项目名称
	Name string `json:"name,omitempty"`
	// ExternalAction 优化目标，枚举值：
	// FOLLOW_ACTION 粉丝增长
	// LIVE_ENGAGEMENT 直播加热
	// LIVE_ENTER_ACTION 直播间观看
	// LIVE_OTO_CLICK 直播间商品点击
	// LIVE_OTO_GROUP_BUYING 直播间团购购买
	// LIVE_STAY_TIME 直播间停留NATIVE_ACTION 用户互动
	// SHOW 展示量
	// OTO_PAY 团购购买
	// POI_RECOMMEND 门店引流
	ExternalAction local.ExternalAction `json:"external_action,omitempty"`
	// DeliveryGoal 投放内容，枚举值：
	// POI 门店
	// PRODUCT 商品
	DeliveryGoal local.DeliveryGoal `json:"delivery_goal,omitempty"`
	// BudgetMode 项目预算类型，枚举值：
	// BUDGET_MODE_DAY 日预算
	// BUDGET_MODE_TOTAL 总预算
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 项目预算，单位为：分
	Budget int64 `json:"budget,omitempty"`
	// BidType 出价方式，枚举值：
	// MANUAL 手动出价
	// SMART 智能出价
	BidType local.BidType `json:"bid_type,omitempty"`
	// Bid 项目出价，单位为：分
	Bid int64 `json:"bid,omitempty"`
	// StartTime 投放开始时间，如：2017-01-01 精确到天
	StartTime string `json:"start_time,omitempty"`
	// EndTime 投放结束时间，如：2017-01-01 精确到天
	EndTime string `json:"end_time,omitempty"`
	// DeliveryPoiMode 本地投放模式
	DeliveryPoiMode local.DeliveryPoiMode `json:"delivery_poi_mode,omitempty"`
	// AutoUpdatePois 自动更新门店开启状态，仅推广全部门店项目返回，枚举值：
	// OFF 不启用
	// ON 启用
	AutoUpdatePois enum.OnOff `json:"auto_update_pois,omitempty"`
	// PoiID 门店id
	// 当项目推广单个门店时返回
	PoiID uint64 `json:"poi_id,omitempty"`
	// MultiPoiNum 推广门店数量
	MultiPoiNum int `json:"multi_poi_num,omitempty"`
	// ProductID 商品ID
	ProductID uint64 `json:"product_id,omitempty"`
	// ScheduleType 投放日期类型 ，枚举值：
	// FROM_NOW_ON 从今天起长期投放
	// START_TO_END 设置开始结束时间
	// FIXED_TIME 固定时长投放
	ScheduleType local.ScheduleType `json:"schedule_type,omitempty"`
	// ScheduleFixedSeconds 投放时长，单位为秒
	ScheduleFixedSeconds int64 `json:"schedule_fixed_seconds,omitempty"`
	// ScheduleTime 投放时段
	// 格式是48*7位字符串，且都是0或1。也就是以半个小时为最小粒度，周一至周日每天分为48个区段，0为不投放，1为投放，不传、全传0、全传1均代表全时段投放。
	// 例如：000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000，则投放时段为周一到周日的11:30~13:30
	ScheduleTime string `json:"schedule_time,omitempty"`
	// IsSetPeakBudget 是否设置峰值预算
	IsSetPeakBudget bool `json:"is_set_peak_budget,omitempty"`
	// HighBudgetRate 上调高峰日预算比例
	// 注意：该字段为百分比，例如：传“40”表示高峰日时预算上调“40%”
	HighBudgetRate int `json:"high_budget_rate,omitempty"`
	// PeakWeekDays 高峰日-自然周
	PeakWeekDays []enum.Weekday `json:"peak_week_days,omitempty"`
	// PeakHolidays 高峰日-节假日
	PeakHolidays []enum.Holiday `json:"peak_holidays,omitempty"`
	// Audience 定向设置
	Audience *Audience `json:"audience,omitempty"`
}

// PoiInfo 项目推广门店信息
type PoiInfo struct {
	// PoiID 门店id
	// 仅推广单门店项目返回，推广多门店时不返回
	PoiID uint64 `json:"poi_id,omitempty"`
	// PoiName 门店名称
	// 仅推广单门店项目返回，推广多门店时不返回
	PoiName string `json:"poi_name,omitempty"`
	// PoiImageURL 门店头图
	// 仅推广单门店项目返回，推广多门店时不返回
	PoiImageURL string `json:"poi_image_url,omitempty"`
	// MultiPoiNum 推广门店数量
	MultiPoiNum int `json:"multi_poi_num,omitempty"`
	// AllPoiMod 推广门店类型，枚举值：
	// ALL 全部门店
	// PART 指定门店
	AllPoiMode local.DeliveryPoiMode `json:"all_poi_mode,omitempty"`
	// AutoUpdatePoi 是否自动更新门店，枚举值:
	AutoUpdatePoi string `json:"auto_update_poi,omitempty"`
}
