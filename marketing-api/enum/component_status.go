package enum

// ComponentStatus 组件审核状态
type ComponentStatus string

const (
	// ComponentStatus_PASS 通过
	ComponentStatus_PASS ComponentStatus = "PASS"
	// ComponentStatus_UNDER 审核中
	ComponentStatus_UNDER ComponentStatus = "UNDER"
	// ComponentStatus_REJECT 未通过
	ComponentStatus_REJECT ComponentStatus = "REJECT"
)
