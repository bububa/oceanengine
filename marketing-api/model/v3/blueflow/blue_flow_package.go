package blueflow

import "github.com/bububa/oceanengine/marketing-api/enum"

// Package 蓝海流量包信息
type Package struct {
	// BlueFlowPackageID 蓝海流量包ID
	BlueFlowPackageID uint64 `json:"blue_flow_package_id,omitempty"`
	// BlueFlowPackageName 蓝海流量包名称
	BlueFlowPackageName string `json:"blue_flow_package_name,omitempty"`
	// FlowCoverageRatio 流量覆盖比例，返回整数（百分比），范围是0% ～ 100%（周维度更新）
	FlowCoverageRatio string `json:"flow_coverage_ratio,omitempty"`
	// EstimatedReachTrafficName 预计触达流量区间（周维度更新）
	EstimatedReachTrafficName []uint64 `json:"estimated_reach_traffic_name,omitempty"`
	// EstimatedCostRange 预估消耗区间（周维度更新）
	EstimatedCostRange []uint64 `json:"estimated_cost_range,omitempty"`
	// SuggestedBidRange 建议出价区间（周维度更新）
	SuggestedBidRange []uint64 `json:"suggested_bid_range,omitempty"`
	// BlueFlowKeyword 流量包蓝海关键词信息
	BlueFlowKeyword []Keyword `json:"blue_flow_keyword,omitempty"`
}

// Keyword 流量包蓝海关键词信息
type Keyword struct {
	// BlueFlowKeywordName 蓝海关键词名称
	BlueFlowKeywordName string `json:"blue_flow_keyword_name,omitempty"`
	// BlueFlowKeywordTopQueryWordList TOP蓝海关键词搜索词（周维度更新）
	BlueFlowKeywordTopQueryWordList []string `json:"blue_flow_keyword_top_query_word_list,omitempty"`
	// BlueFlowKeywordEstimatedCostRange 蓝海关键词预估消耗区间（周维度更新）
	BlueFlowKeywordEstimatedCostRange []uint64 `json:"blue_flow_keyword_estimated_cost_range,omitempty"`
	// SupportBlueFlowRecommendMaterial 是否支持设置系统推荐图文（周维度更新）：
	// 是 true
	// 否false
	SupportBlueFlowRecommendMaterial bool `json:"support_blue_flow_recommend_material,omitempty"`
	// BlueFlowRecommendType 关键词推荐类别
	// 枚举值: ORDINARY普通关键词，PREFERRED优选关键词
	BlueFlowRecommendType enum.BlueFlowRecommendType `json:"blue_flow_recommend_type,omitempty"`
}
