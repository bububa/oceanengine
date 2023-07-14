package enum

// WechatAuditStatus 微信小程序审核状态
type WechatAuditStatus string

const (
	// WechatAuditStatus_AUDIT_ACCEPTED 审核通过
	WechatAuditStatus_AUDIT_ACCEPTED WechatAuditStatus = "AUDIT_ACCEPTED"
	// WechatAuditStatus_AUDITING 审核中
	WechatAuditStatus_AUDITING WechatAuditStatus = "AUDITING"
	// WechatAuditStatus_AUDIT_REJECTED 审核拒绝
	WechatAuditStatus_AUDIT_REJECTED WechatAuditStatus = "AUDIT_REJECTED"
)
