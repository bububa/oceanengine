package enum

// FlowPackageStatus 流量包状态
type FlowPackageStatus string

const (
	// FLOW_PACKAGE_ENABLE已启用
	FLOW_PACKAGE_ENABLE FlowPackageStatus = "FLOW_PACKAGE_ENABLE"
	// FLOW_PACKAGE_DISABLE 已废弃
	FLOW_PACKAGE_DISABLE FlowPackageStatus = "FLOW_PACKAGE_DISABLE"
)
