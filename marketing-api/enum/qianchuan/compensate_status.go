package qianchuan

// CompensateRequestStatus 当前请求是否成功
type CompensateRequestStatus string

const (
	// CompensateRequestStatus_SUCCESS 查询成功
	CompensateRequestStatus_SUCCESS CompensateRequestStatus = "SUCCESS"
	// CompensateRequestStatus_FAILED 查询失败，请重试
	CompensateRequestStatus_FAILED CompensateRequestStatus = "FAILED"
)

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
)
