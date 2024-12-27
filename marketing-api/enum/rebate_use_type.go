package enum

// RebateUseType 使用类型
type RebateUseType string

const (
	// RebateUseType_CASH 现金
	RebateUseType_CASH RebateUseType = "CASH"
	// RebateUseType_CHARGE 充值
	RebateUseType_CHARGE RebateUseType = "CHARGE"
	// RebateUseType_HEDGING 抵扣
	RebateUseType_HEDGING RebateUseType = "HEDGING"
	// RebateUseType_NONPAYMENT 无需支付
	RebateUseType_NONPAYMENT RebateUseType = "NONPAYMENT"
	// RebateUseType_REVERT 负数冲抵
	RebateUseType_REVERT RebateUseType = "REVERT"
)
