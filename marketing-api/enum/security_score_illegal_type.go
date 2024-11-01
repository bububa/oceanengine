package enum

// SecurityScoreIllegalType 违规类型
type SecurityScoreIllegalType string

const (
	// SecurityScoreIllegalType_GENERAL 一般违规（AD）
	SecurityScoreIllegalType_GENERAL SecurityScoreIllegalType = "GENERAL"
	// SecurityScoreIllegalType_SERIOUS 严重违规（AD）
	SecurityScoreIllegalType_SERIOUS SecurityScoreIllegalType = "SERIOUS"
)
