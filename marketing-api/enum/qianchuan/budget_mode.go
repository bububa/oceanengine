package qianchuan

// BudgetMode 预算类型
type BudgetMode string

const (
	// BudgetMode_INFINITE 不限
	BudgetMode_INFINITE BudgetMode = "INFINITE"
	// BudgetMode_SPECIFIED 日预算
	BudgetMode_SPECIFIED BudgetMode = "SPECIFIED"
)
