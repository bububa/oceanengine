package enum

// TaskRaiseBudgetMode 优选起量任务预算设置
type TaskRaiseBudgetMode string

const (
	// TaskRaiseBudgetMode_LIMIT 指定预算
	TaskRaiseBudgetMode_LIMIT TaskRaiseBudgetMode = "LIMIT"
	// TaskRaiseBudgetMode_NO_LIMIT 不限预算
	TaskRaiseBudgetMode_NO_LIMIT TaskRaiseBudgetMode = "NO_LIMIT"
)
