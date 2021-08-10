package enum

// SiteAuditStatus 落地页组站点审核状态
type SiteAuditStatus string

const (
	// SiteAuditStatus_AUDIT_UNKNOWN 未知中间态
	SiteAuditStatus_AUDIT_UNKNOWN SiteAuditStatus = "AUDIT_UNKNOWN"
	// SiteAuditStatus_AUDIT_ACCEPTED 审核通过
	SiteAuditStatus_AUDIT_ACCEPTED SiteAuditStatus = "AUDIT_ACCEPTED"
	// SiteAuditStatus_AUDIT_REJECTED 审核拒绝
	SiteAuditStatus_AUDIT_REJECTED SiteAuditStatus = "AUDIT_REJECTED"
	// SiteAuditStatus_BANNED 审核封禁
	SiteAuditStatus_BANNED SiteAuditStatus = "BANNED"
	// SiteAuditStatus_AUDITING 审核中
	SiteAuditStatus_AUDITING SiteAuditStatus = "AUDITING"
	// SiteAuditStatus_AWAIT_AUDIT 待审核
	SiteAuditStatus_AWAIT_AUDIT SiteAuditStatus = "AWAIT_AUDIT"
)
