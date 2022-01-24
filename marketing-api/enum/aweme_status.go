package enum

// AwemeStatus 抖音号带货状态
type AwemeStatus string

const (
	// AwemeStatus_NORMAL 可以正常投放
	AwemeStatus_NORMAL AwemeStatus = "NORMAL"
	// AwemeStatus_ANCHOR_FORBID 带货口碑分过低，暂时无法创建计划
	AwemeStatus_ANCHOR_FORBID AwemeStatus = "ANCHOR_FORBID"
	// AwemeStatus_ANCHOR_REACH_UPPER_LIMIT_TODAY 带货分过低或暂无带货分，可以创建计划，但无法产生消耗，带货分恢复正常后可正常消耗
	AwemeStatus_ANCHOR_REACH_UPPER_LIMIT_TODAY AwemeStatus = "ANCHOR_REACH_UPPER_LIMIT_TODAY"
)
