package enum

// ActivateType 新用户使用时间
type ActivateType string

const (
	// ActivateType_WITH_IN_A_MONTH 一个月以内
	ActivateType_WITH_IN_A_MONTH ActivateType = "WITH_IN_A_MONTH"
	// ActivateType_ONE_MONTH_2_THREE_MONTH 一个月到三个月
	ActivateType_ONE_MONTH_2_THREE_MONTH ActivateType = "ONE_MONTH_2_THREE_MONTH"
	// ActivateType_THREE_MONTH_EAILIER 三个月或更早
	ActivateType_THREE_MONTH_EAILIER ActivateType = "THREE_MONTH_EAILIER"
)
