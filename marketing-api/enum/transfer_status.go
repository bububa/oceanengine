package enum

// TransferStatus 转账总状态
type TransferStatus string

const (
	// TransferStatus_NO_TRANSFER 未转账
	TransferStatus_NO_TRANSFER TransferStatus = "NO_TRANSFER"
	// TransferStatus_TRANSFER_FAILED 转账失败(终态)
	TransferStatus_TRANSFER_FAILED TransferStatus = "TRANSFER_FAILED"
	// TransferStatus_TRANSFER_ING 转账中
	TransferStatus_TRANSFER_ING TransferStatus = "TRANSFER_ING"
	// TransferStatus_TRANSFER_PART 转账部分成功(终态)
	TransferStatus_TRANSFER_PART TransferStatus = "TRANSFER_PART"
	// TransferStatus_TRANSFER_SUCCESS 转账成功(终态)
	TransferStatus_TRANSFER_SUCCESS TransferStatus = "TRANSFER_SUCCESS"
)
