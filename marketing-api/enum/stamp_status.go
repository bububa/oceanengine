package enum

// StampStatus 盖章状态
type StampStatus string

const (
	// StampStatus_NO 未盖章
	StampStatus_NO StampStatus = "NO"
	// StampStatus_STAMPING 审批中
	StampStatus_STAMPING StampStatus = "STAMPING"
	// StampStatus_STAMPED 已盖章
	StampStatus_STAMPED StampStatus = "STAMPED"
)
