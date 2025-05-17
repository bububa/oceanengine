package enum

// TransferDirection 转账方向，以目标账户视角确定
type TransferDirection string

const (
	// TRANSFER_IN 转入
	TRANSFER_IN TransferDirection = "TRANSFER_IN"
	// TRANSFER_OUT 转出
	TRANSFER_OUT TransferDirection = "TRANSFER_OUT"
)
