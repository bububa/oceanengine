package qianchuan

// AdOptStatus 广告计划操作状态
type AdOptStatus string

const (
	// AdOptStatus_ENABLE 启用
	AdOptStatus_ENABLE AdOptStatus = "ENABLE"
	// AdOptStatus_DISABLE 暂停
	AdOptStatus_DISABLE AdOptStatus = "DISABLE"
	// AdOptStatus_DELETE 删除
	AdOptStatus_DELETE AdOptStatus = "DELETE"
)
