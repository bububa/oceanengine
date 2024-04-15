package enum

// BudgetGroupStatus 预算组聚合状态
type BudgetGroupStatus string

const (
	// BudgetGroupStatus_DELETED 已删除
	BudgetGroupStatus_DELETED BudgetGroupStatus = "DELETED"
	// BudgetGroupStatus_ENABLE 启用中
	BudgetGroupStatus_ENABLE BudgetGroupStatus = "ENABLE"
	// BudgetGroupStatus_UNDELIVERIED 未投放
	BudgetGroupStatus_UNDELIVERIED BudgetGroupStatus = "UNDELIVERIED"
	// BudgetGroupStatus_ACCOUNT_EXCEEDED 账户超出预算
	BudgetGroupStatus_ACCOUNT_EXCEEDED BudgetGroupStatus = "ACCOUNT_EXCEEDED"
	// BudgetGroupStatus_GROUP_EXCEEDED 预算组超出预算
	BudgetGroupStatus_GROUP_EXCEEDED BudgetGroupStatus = "GROUP_EXCEEDED"
)
