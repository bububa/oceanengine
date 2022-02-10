package enum

// PricingType 计划出价类型
type PricingType string

const (
	// PRICING_CPC CPC（点击付费），出价范围（单位元）:0.2-100; 日预算范围（单位元）：大于100;总预算范围：大于最低日预算乘投放天数
	PRICING_CPC PricingType = "PRICING_CPC"
	// PRICING_CPM CPM（展示付费），出价范围（单位元）: 4-100; 日预算范围（单位元）：大于100; 总预算范围：大于最低日预算乘投放天数
	PRICING_CPM PricingType = "PRICING_CPM"
	// PRICING_OCPC OCPC（已下线，仅投放范围为穿山甲可用）
	PRICING_OCPC PricingType = "PRICING_OCPC"
	// PRICING_OCPM OCPM（转化量付费），出价范围（单位元）:0.1-10000; 日预算范围（单位元）：大于300;总预算范围：大于最低日预算乘投放天数
	PRICING_OCPM PricingType = "PRICING_OCPM"
	// PRICING_CPA CPV ，出价范围（单位元）:0.07-100; 日预算范围（单位元）：大于100; 总预算范围：大于最低日预算乘投放天数; （CPV广告只支持投放到头条系广告位，不支持投放到视频信息流如西瓜、火山、抖音）
	PRICING_CPV PricingType = "PRICING_CPV"
	// PRICING_CPA CPA（已下线，不支持进行出价方式为CPA的投放行为）
	PRICING_CPA PricingType = "PRICING_CPA"
)
