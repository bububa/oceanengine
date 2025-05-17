package enum

// EventConvertOptimizedGoalValueType 设置差异价值
type EventConvertOptimizedGoalValueType = string

const (
	// EventConvertOptimizedGoalValueType_DISABLED 未设置
	EventConvertOptimizedGoalValueType_DISABLED EventConvertOptimizedGoalValueType = "DISABLED"
	// EventConvertOptimizedGoalValueType_DISCRIMINATEBYGROUP 人群差异价值
	EventConvertOptimizedGoalValueType_DISCRIMINATEBYGROUP EventConvertOptimizedGoalValueType = "DISCRIMINATEBYGROUP"
	// EventConvertOptimizedGoalValueType_DYNAMICVALUE 动态价值
	EventConvertOptimizedGoalValueType_DYNAMICVALUE EventConvertOptimizedGoalValueType = "DYNAMICVALUE"
	// EventConvertOptimizedGoalValueType_FIXED 固定价值
	EventConvertOptimizedGoalValueType_FIXED EventConvertOptimizedGoalValueType = "FIXED"
)
