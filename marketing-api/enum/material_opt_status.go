package enum

// MaterialOptStatus 素材操作状态
type MaterialOptStatus string

const (
	// MaterialOptStatus_ENABLE 启用
	MaterialOptStatus_ENABLE MaterialOptStatus = "ENABLE"
	// MaterialOptStatus_DISABLE 暂停
	MaterialOptStatus_DISABLE MaterialOptStatus = "DISABLE"
)
