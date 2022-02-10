package enum

// TransactionType 流水类型
type TransactionType string

const (
	// RECHARGE 充值
	RECHARGE TransactionType = "RECHARGE"
	// TRANSFER 转账
	TRANSFER TransactionType = "TRANSFER"
)
