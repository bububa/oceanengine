package enum

// AIRepairStatus 采纳状态
type AIRepairStatus string

const (
	// AIRepairStatus_SUCCESS 采纳成功
	AIRepairStatus_SUCCESS AIRepairStatus = "SUCCESS"
	// AIRepairStatus_FAILED  采纳失败
	AIRepairStatus_FAILED AIRepairStatus = "FAILED"
	// AIRepairStatus_PROCESSING 采纳中
	AIRepairStatus_PROCESSING AIRepairStatus = "PROCESSING"
	// AIRepairStatus_NO_TASK  采纳任务不存在，可能为未创建采纳任务/已过期/广告已删除/广告下素材已删除
	AIRepairStatus_NO_TASK AIRepairStatus = "NO_TASK"
)
