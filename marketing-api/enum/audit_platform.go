package enum

// AuditPlatform 审核来源类型
type AuditPlatform string

const (
	// AuditPlatform_AD 广告审核
	AuditPlatform_AD AuditPlatform = "AD"
	// AuditPlatform_CONTENT 内容审核
	AuditPlatform_CONTENT AuditPlatform = "CONTENT"
)
