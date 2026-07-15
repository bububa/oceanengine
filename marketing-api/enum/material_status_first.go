package enum

// MaterialStatusFirst 素材一级状态
type MaterialStatusFirst string

const (
	// MaterialStatusFirst_STATUS_DELETE 已删除
	MaterialStatusFirst_STATUS_DELETE MaterialStatusFirst = "STATUS_DELETE"
	// MaterialStatusFirst_STATUS_DISABLE 未投放
	MaterialStatusFirst_STATUS_DISABLE MaterialStatusFirst = "STATUS_DISABLE"
	// MaterialStatusFirst_STATUS_FROZEN 已终止
	MaterialStatusFirst_STATUS_FROZEN MaterialStatusFirst = "STATUS_FROZEN"
	// MaterialStatusFirst_STATUS_OK 投放中
	MaterialStatusFirst_STATUS_OK MaterialStatusFirst = "STATUS_OK"
	// MaterialStatusFirst_STATUS_TIME_DONE 已完成
	MaterialStatusFirst_STATUS_TIME_DONE MaterialStatusFirst = "STATUS_TIME_DONE"
)
