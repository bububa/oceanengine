package enum

// ProductAuditStatus 商品审核状态
type ProductAuditStatus string

const (
	// ProductAuditStatus_AUDIT_STATUS_APPROVE 审核通过
	ProductAuditStatus_AUDIT_STATUS_APPROVE ProductAuditStatus = "AUDIT_STATUS_APPROVE"
	// ProductAuditStatus_AUDIT_STATUS_INIT 审核中
	ProductAuditStatus_AUDIT_STATUS_INIT ProductAuditStatus = "AUDIT_STATUS_INIT"
	// ProductAuditStatus_AUDIT_STATUS_REJECT 审核未通过
	ProductAuditStatus_AUDIT_STATUS_REJECT ProductAuditStatus = "AUDIT_STATUS_REJECT"
)
