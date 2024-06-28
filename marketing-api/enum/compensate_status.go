package enum

// CompensateStatus 计划成本保障状态
type CompensateStatus string

const (
	// CompensateStatus_IN_EFFECT 成本保障生效中
	CompensateStatus_IN_EFFECT CompensateStatus = "IN_EFFECT"
	// CompensateStatus_INVALID 成本保障已失效
	CompensateStatus_INVALID CompensateStatus = "INVALID"
	// CompensateStatus_CONFIRMING 成本保障确认中
	CompensateStatus_CONFIRMING CompensateStatus = "CONFIRMING"
	// CompensateStatus_PAID 成本保障已到账
	CompensateStatus_PAID CompensateStatus = "PAID"
	// CompensateStatus_ENDED 成本保障已结束
	CompensateStatus_ENDED CompensateStatus = "ENDED"
	// CompensateStatus_DEFAULT 无成本保障状态
	CompensateStatus_DEFAULT CompensateStatus = "DEFAULT"
	// CompensateStatus_FAILED 成本保障查询超时，请重试
	CompensateStatus_FAILED CompensateStatus = "FAILED"
	// CompensateStatus_UNSUPPORTED 查询失败，传入参数project信息有误，多为project_id在账户下不存在
	CompensateStatus_UNSUPPORTED CompensateStatus = "UNSUPPORTED"
)
