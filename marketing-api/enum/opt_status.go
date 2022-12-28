package enum

// OptStatus 目标操作
type OptStatus string

const (
	// OptStatus_ENABLE 启用项目
	OptStatus_ENABLE OptStatus = "ENABLE"
	// OptStatus_DISABLE暂停项目
	OptStatus_DISABLE OptStatus = "DISABLE"
)
