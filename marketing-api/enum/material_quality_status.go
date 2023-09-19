package enum

// MaterialQualityStatus 连山投前分析状态
type MaterialQualityStatus string

const (
	// MaterialQualityStatus_AUDITING 审核中
	MaterialQualityStatus_AUDITING MaterialQualityStatus = "AUDITING"
	// MaterialQualityStatus_AUDIT_ACCEPTED 预审通过
	MaterialQualityStatus_AUDIT_ACCEPTED MaterialQualityStatus = "AUDIT_ACCEPTED"
	// MaterialQualityStatus_AUDIT_FAILED 预审失败
	MaterialQualityStatus_AUDIT_FAILED MaterialQualityStatus = "AUDIT_FAILED"
	// MaterialQualityStatus_AUDIT_REJECT 预审拒绝
	MaterialQualityStatus_AUDIT_REJECT MaterialQualityStatus = "AUDIT_REJECT"
	// MaterialQualityStatus_AUDIT_SUBMIT 发起预审
	MaterialQualityStatus_AUDIT_SUBMIT MaterialQualityStatus = "AUDIT_SUBMIT"
	// MaterialQualityStatus_AUDIT_TIMEOUT 预审超时
	MaterialQualityStatus_AUDIT_TIMEOUT MaterialQualityStatus = "AUDIT_TIMEOUT"
)
