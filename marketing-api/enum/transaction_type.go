package enum

type TransactionType string

const (
	RECHARGE TransactionType = "RECHARGE" // 充值
	TRANSFER TransactionType = "TRANSFER" // 转账
)
