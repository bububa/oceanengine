package enum

// QuickAppStatus 应用状态
type QuickAppStatus string

const (
	// QuickAppStatus_AUDIT_DOING 审核中
	QuickAppStatus_AUDIT_DOING QuickAppStatus = "AUDIT_DOING"
	// QuickAppStatus_AUDIT_SEND_REJECTED 送审失败
	QuickAppStatus_AUDIT_SEND_REJECTED QuickAppStatus = "AUDIT_SEND_REJECTED"
	// QuickAppStatus_AUDIT_REJECTED 审核失败
	QuickAppStatus_AUDIT_REJECTED QuickAppStatus = "AUDIT_REJECTED"
	// QuickAppStatus_AUDIT_ACCEPTED 审核成功
	QuickAppStatus_AUDIT_ACCEPTED QuickAppStatus = "AUDIT_ACCEPTED"
	// QuickAppStatus_REMOVED 已下架
	QuickAppStatus_REMOVED QuickAppStatus = "REMOVED"
)
