package enum

// AndroidBasicPackageStatus 状态，枚举值：
type AndroidBasicPackageStatus string

const (
	// AndroidBasicPackageStatus_ALL 不限
	AndroidBasicPackageStatus_ALL AndroidBasicPackageStatus = "ALL"
	// AndroidBasicPackageStatus_AUDIT_ACCEPTED 审核通过
	AndroidBasicPackageStatus_AUDIT_ACCEPTED AndroidBasicPackageStatus = "AUDIT_ACCEPTED"
	// AndroidBasicPackageStatus_AUDIT_DOING 审核中
	AndroidBasicPackageStatus_AUDIT_DOING AndroidBasicPackageStatus = "AUDIT_DOING"
	// AndroidBasicPackageStatus_AUDIT_REJECTED 审核拒绝
	AndroidBasicPackageStatus_AUDIT_REJECTED AndroidBasicPackageStatus = "AUDIT_REJECTED"
	// AndroidBasicPackageStatus_ENABLE 可用
	AndroidBasicPackageStatus_ENABLE AndroidBasicPackageStatus = "ENABLE"
)
