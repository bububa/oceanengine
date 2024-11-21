package enum

// StarAdUniteTaskAuditStatus 任务审核状态 可选值:
type StarAdUniteTaskAuditStatus string

const (
	// StarAdUniteTaskAuditStatus_CONFIRM 审核通过
	StarAdUniteTaskAuditStatus_CONFIRM StarAdUniteTaskAuditStatus = "CONFIRM"
	// StarAdUniteTaskAuditStatus_CONFIRM_FAIL 审核失败
	StarAdUniteTaskAuditStatus_CONFIRM_FAIL StarAdUniteTaskAuditStatus = "CONFIRM_FAIL"
	// StarAdUniteTaskAuditStatus_PENDING_CONFIRM 审核中
	StarAdUniteTaskAuditStatus_PENDING_CONFIRM StarAdUniteTaskAuditStatus = "PENDING_CONFIRM"
)
