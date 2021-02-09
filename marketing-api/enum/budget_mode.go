package enum

type BudgetMode string

const (
	BUDGET_MODE_INFINITE BudgetMode = "BUDGET_MODE_INFINITE" // 不限
	BUDGET_MODE_DAY      BudgetMode = "BUDGET_MODE_DAY"      // 日预算
	BUDGET_MODE_TOTAL    BudgetMode = "BUDGET_MODE_TOTAL"    // 总预算
)
