package enum

// FlowControlMode  竞价策略
type FlowControlMode string

const (
	// FLOW_CONTROL_MODE_FAST 优先跑量（对应CPC的加速投放
	FLOW_CONTROL_MODE_FAST FlowControlMode = "FLOW_CONTROL_MODE_FAST"
	// FLOW_CONTROL_MODE_SMOOTH 优先低成本（对应CPC的标准投放）
	FLOW_CONTROL_MODE_SMOOTH FlowControlMode = "FLOW_CONTROL_MODE_SMOOTH"
	// FLOW_CONTROL_MODE_BALANCE 均衡投放（新增字段）
	FLOW_CONTROL_MODE_BALANCE FlowControlMode = "FLOW_CONTROL_MODE_BALANCE"
)
