package enum

// MicroAppAuditStatus 审核状态
type MicroAppAuditStatus string

const (
	// MicroAppAuditStatus_AUDIT_ACCEPTED 审核通过
	MicroAppAuditStatus_AUDIT_ACCEPTED MicroAppAuditStatus = "AUDIT_ACCEPTED"
	// MicroAppAuditStatus_AUDITING 审核中
	MicroAppAuditStatus_AUDITING MicroAppAuditStatus = "AUDITING"
	// MicroAppAuditStatus_AUDIT_REJECTED 审核不通过
	MicroAppAuditStatus_AUDIT_REJECTED MicroAppAuditStatus = "AUDIT_REJECTED"
	// MicroAppAuditStatus_ALL 全部（默认值）
	MicroAppAuditStatus_ALL MicroAppAuditStatus = "ALL"
)
