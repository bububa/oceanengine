package enum

// StarAdUniteTaskEvaluateType 任务优化目标 可选值:
type StarAdUniteTaskEvaluateType string

const (
	// StarAdUniteTaskEvaluateType_ACTIVE 激活
	StarAdUniteTaskEvaluateType_ACTIVE StarAdUniteTaskEvaluateType = "ACTIVE"
	// StarAdUniteTaskEvaluateType_ACTIVE_PAY 首次付费
	StarAdUniteTaskEvaluateType_ACTIVE_PAY StarAdUniteTaskEvaluateType = "ACTIVE_PAY"
	// StarAdUniteTaskEvaluateType_DEEP_PURCHASE 每次付费
	StarAdUniteTaskEvaluateType_DEEP_PURCHASE StarAdUniteTaskEvaluateType = "DEEP_PURCHASE"
	// StarAdUniteTaskEvaluateType_INSTALL_FINISH 安装完成
	StarAdUniteTaskEvaluateType_INSTALL_FINISH StarAdUniteTaskEvaluateType = "INSTALL_FINISH"
	// StarAdUniteTaskEvaluateType_REGISTER 注册
	StarAdUniteTaskEvaluateType_REGISTER StarAdUniteTaskEvaluateType = "REGISTER"
)
