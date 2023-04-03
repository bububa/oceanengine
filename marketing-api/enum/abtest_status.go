package enum

// AbTestStatus 根据实验状态过滤，允许值：，，，"FAILED"：不满足实验条件。
type AbTestStatus string

const (
	// AbTestStatus_CREATED 排期中
	AbTestStatus_CREATED AbTestStatus = "CREATED"
	// AbTestStatus_PROCESSING 进行中
	AbTestStatus_PROCESSING AbTestStatus = "PROCESSING"
	// AbTestStatus_FINISH 结束
	AbTestStatus_FINISH AbTestStatus = "FINISH"
	// AbTestStatus_FAILED 不满足实验条件
	AbTestStatus_FAILED AbTestStatus = "FAILED"
)
