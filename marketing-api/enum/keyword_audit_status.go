package enum

// KeywordAuditStatus  审核状态
type KeywordAuditStatus string

const (
	// KeywordAuditStatus_AUDITING 审核中
	KeywordAuditStatus_AUDITING KeywordAuditStatus = "AUDITING"
	// KeywordAuditStatus_AUDIT_ACCEPTED 审核通过
	KeywordAuditStatus_AUDIT_ACCEPTED KeywordAuditStatus = "AUDIT_ACCEPTED"
	// KeywordAuditStatus_AUDIT_REJECTED 审核拒绝
	KeywordAuditStatus_AUDIT_REJECTED KeywordAuditStatus = "AUDIT_REJECTED"
	// KeywordAuditStatus_DELETED 已删除
	KeywordAuditStatus_DELETED KeywordAuditStatus = "DELETED"
)
