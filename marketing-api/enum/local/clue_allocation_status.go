package local

// ClueAllocationStatus  分配状态
type ClueAllocationStatus string

const (
	// ClueAllocationStatus_NOT_ASSIGN 待分配
	ClueAllocationStatus_NOT_ASSIGN ClueAllocationStatus = "NOT_ASSIGN"
	// ClueAllocationStatus_ASSIGNED 已分配
	ClueAllocationStatus_ASSIGNED ClueAllocationStatus = "ASSIGNED"
)
