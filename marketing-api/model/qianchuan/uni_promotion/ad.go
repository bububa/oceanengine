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
	// StatsInfo 消耗指标
	StatsInfo *report.UniPromotionStats `json:"stats_info,omitempty"`
	// RoomInfo 主播信息
	RoomInfo []live.Room `json:"room_info,omitempty"`
	// ProductInfo 商品信息
	ProductInfo []ProductInfo `json:"product_info,omitempty"`
}

// AdInfo 广告信息
type AdInfo struct {
	// ID 推广id
	ID uint64 `json:"id,omitempty"`
	// Name 投放名称
	Name string `json:"name,omitempty"`
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
	// BudgetMode  预算类型
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Status 投放状态
	Status qianchuan.AdStatus `json:"status,omitempty"`
	// OptStatus 操作状态，详见【附录-枚举值】
	OptStatus qianchuan.AdOptStatus `json:"opt_status,omitempty"`
	// SmartBidType  投放方式，可选值:
	// SMART_BID_CONSERVATIVE 放量投放
	// SMART_BID_CUSTOM 控成本投放
	SmartBidType enum.SmartBidType `json:"smart_bid_type,omitempty"`
	// Roi2Goal 支付ROI目标，最多支持两位小数
	Roi2Goal float64 `json:"roi2_goal,omitempty"`
	// Budget 预算
	Budget float64 `json:"budget,omitempty"`
	// DeliverySecond 投放时长
	DeliverySecond int64 `json:"delivery_second,omitempty"`
	// DailyDeliveryTime 每日投放时长，0.5～24h
	DailyDeliveryTime float64 `json:"daily_delivery_time,omitempty"`
	// CompensateInfo 成本保障信息
	CompensateInfo *CompensateInfo `json:"compensate_info,omitempty"`
	// DeepExternalAction 优化目标，可选值:
	// AD_CONVERT_TYPE_LIVE_PAY_ROI 成交ROI
	// AD_CONVERT_TYPE_LIVE_PURE_PAY_ROI 净成交ROI
	DeepExternalAction qianchuan.ExternalAction `json:"deep_external_action,omitempty"`
}

// CompensateInfo 成本保障信息
type CompensateInfo struct {
	// Status 当前请求是否成功，可选值:
	// SUCCESS 查询成功
	// FAILED 查询失败
	Status qianchuan.CompensateRequestStatus `json:"status,omitempty"`
	//  计划成本保障状态，可选值:
	// IN_EFFECT 成本保障生效中
	// INVALID 成本保障已失效
	// CONFIRMING 成本保障确认中
	// PAID 成本保障已到账
	// ENDED 成本保障已结束
	// DEFAULT 无成本保障状态
	CompensateStatus qianchuan.CompensateStatus `json:"compensate_status,omitempty"`
}

// ProductInfo 商品信息
type ProductInfo struct {
	// ProductID 商品ID
	ProductID uint64 `json:"product_id,omitempty"`
	// ProductName 商品名称
	ProductName string `json:"product_name,omitempty"`
	// ProductImage 商品预览图
	ProductImage string `json:"product_image,omitempty"`
	// RecommendReasons 推荐理由
	RecommendReasons []string `json:"recommend_reasons,omitempty"`
	// ChannelType 渠道类型，可选值:
	// SHOP_SELL 商家自卖
	// STAR_SELL 达人自播
	ChannelType qianchuan.ProductChannelType `json:"channel_type,omitempty"`
	// ChannelID 渠道id，渠道品相关介绍见《【抖店】销售渠道品功能操作手册》
	ChannelID uint64 `json:"channel_id,omitempty"`
}

// AdDetail 广告详情
type AdDetail struct {
	// ProductInfos 商品信息
	ProductInfos []ProductInfo `json:"product_infos,omitempty"`
	// DeliverySetting 投放设置
	DeliverySetting *DeliverySetting `json:"delivery_setting,omitempty"`
	// CreativeSetting 创意设置
	CreativeSetting *CreativeSetting `json:"creative_setting,omitempty"`
	// RoomInfo 主播信息
	RoomInfo []live.Room `json:"room_info,omitempty"`
	// ProgrammaticCreativeMediaList 程序化创意信息
	ProgrammaticCreativeMediaList []ProgrammaticCreativeMedia `json:"programmatic_creative_media_list,omitempty"`
	// MultiProductCreativeList 商品全域创意素材信息
	// 支持0-100个素材+0-30个标题+1个投放卡片
	// 注意：对于无号商家，仅支持投放商品卡体裁
	MultiProductCreativeList []MultiProductCreative `json:"multi_product_creative_list,omitempty"`
	// Name 推广名称
	Name string `json:"name,omitempty"`
	// ModifyTime 修改时间
	ModifyTime string `json:"modify_time,omitempty"`
	// CreateTime 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// MarketingGoal 营销目标
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// Status 投放状态
	Status qianchuan.AdStatus `json:"status,omitempty"`
	// OptStatus 操作状态，详见【附录-枚举值】
	OptStatus qianchuan.AdOptStatus `json:"opt_status,omitempty"`
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// AwemeID 抖音号id
	AwemeID uint64 `json:"aweme_id,omitempty"`
}
