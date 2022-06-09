package enum

// AdOptStatus 广告计划操作类型
type AdOptStatus string

const (
	// AdOptStatus_ENABLE 启用
	AdOptStatus_ENABLE AdOptStatus = "AD_STATUS_ENABLE"
	// AdOptStatus_DISABLE 暂停
	AdOptStatus_DISABLE AdOptStatus = "AD_STATUS_DISABLE"
	// AdOptStatus_DISABLE_BY_QUOTA 当前账户的在投广告配额达限、因此系统暂停该广告 （若您账户下查询不到该状态的计划，表明您账户下的计划未受影响）
	AdOptStatus_DISABLE_BY_QUOTA AdOptStatus = "AD_STATUS_DISABLE_BY_QUOTA"
)
