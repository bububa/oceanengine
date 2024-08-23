package eventmanager

import "github.com/bububa/oceanengine/marketing-api/enum"

// EventConvertOptimizedGoal 优化目标数据
type EventConvertOptimizedGoal struct {
	// ExternalAction 预定义转化目标，具体枚举可查看【附录-预定义转化类型】
	ExternalAction enum.AdConvertType `json:"external_action,omitempty"`
	// OptimizationName 事件名称
	OptimizationName string `json:"optimization_name,omitempty"`
	// HistoryBack 历史有无回传, true 表示有，false 表示无
	HistoryBack bool `json:"history_back,omitempty"`
	// TwentyFourHourBack 24 小时历史有无回传, true 表示有，false 表示无
	TwentyFourHourBack bool `json:"twenty_four_hour_back,omitempty"`
	// ValueType 价值类型，Disabled 不展示、DiscriminateByGroup 人群差异价值、DynamicValue 动态回传价值、Fixed 固定价值
	ValueType enum.EventConvertOptimizedGoalValueType `json:"value_type,omitempty"`
	// AssetTypes 资产类型，:THIRD_EXTERNAL:三方落地页、TETRIS_EXTERNAL:建站、APP 应用、QUICK_APP 快应用、MINI_PROGRAME字节小程序
	AssetTypes []enum.AssetType `json:"asset_types,omitempty"`
	// DeepGoals 深度优化目标列表
	DeepGoals []EventConvertOptimizedGoal `json:"deep_goals,omitempty"`
	// DeepExternalAction 深度转化目标，具体枚举可查看【附录-预定义转化类型】
	DeepExternalAction enum.DeepExternalAction `json:"deep_external_action,omitempty"`
}
