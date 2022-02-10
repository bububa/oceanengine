package enum

// AdOptStatus 广告计划操作类型
type AdOptStatus string

const (
	// AdOptStatus_ENABLE 启用
	AdOptStatus_ENABLE AdOptStatus = "AD_STATUS_ENABLE"
	// AdOptStatus_DISABLE 暂停
	AdOptStatus_DISABLE AdOptStatus = "AD_STATUS_DISABLE"
)
