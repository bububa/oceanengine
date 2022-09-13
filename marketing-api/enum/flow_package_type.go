package enum

// FlowPackageType 按照流量包类型过滤，允许值：
type FlowPackageType string

const (
	// FlowPackageType_CUSTOMIZE 自定义
	FlowPackageType_CUSTOMIZE FlowPackageType = "CUSTOMIZE"
	// FlowPackageType_FEATURED 运营推荐
	FlowPackageType_FEATURED FlowPackageType = "FEATURED"
	// FlowPackageType_SYSTEM 系统推荐
	FlowPackageType_SYSTEM FlowPackageType = "SYSTEM"
)
