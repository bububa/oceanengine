package enum

// SecurityScoreViolationEventStatus 生效状态
type SecurityScoreViolationEventStatus string

const (
	// SecurityScoreViolationEventStatus_APPEAL 已申诉（失效）
	SecurityScoreViolationEventStatus_APPEAL SecurityScoreViolationEventStatus = "APPEAL"
	// SecurityScoreViolationEventStatus_FAILAPPEAL 申诉失败
	SecurityScoreViolationEventStatus_FAILAPPEAL SecurityScoreViolationEventStatus = "FAILAPPEAL"
	// SecurityScoreViolationEventStatus_ONAPPEAL 申诉中
	SecurityScoreViolationEventStatus_ONAPPEAL SecurityScoreViolationEventStatus = "ONAPPEAL"
	// SecurityScoreViolationEventStatus_VALID 生效
	SecurityScoreViolationEventStatus_VALID SecurityScoreViolationEventStatus = "VALID"
)
