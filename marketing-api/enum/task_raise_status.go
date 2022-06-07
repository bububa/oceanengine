package enum

// TaskRaiseStatus 优选起量任务状态
type TaskRaiseStatus string

const (
	// TaskRaiseStatus_UNKNOWN
	TaskRaiseStatus_UNKNOWN TaskRaiseStatus = ""
	// TaskRaiseStatus_RAISING 起量中`
	TaskRaiseStatus_RAISING TaskRaiseStatus = "RAISING"
	// TaskRaiseStatus_STOP 已结束
	TaskRaiseStatus_STOP TaskRaiseStatus = "STOP"
	// TaskRaiseStatus_DISABLERAISE 不能开启起量
	TaskRaiseStatus_DISABLERAISE TaskRaiseStatus = "DISABLERAISE"
	// TaskRaiseStatus_ENABLERAISE 可以开启起量
	TaskRaiseStatus_ENABLERAISE TaskRaiseStatus = "ENABLERAISE"
)
