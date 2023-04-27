package enum

// VideoMaterialClearTaskStatus 视频素材清理任务状态
type VideoMaterialClearTaskStatus string

const (
	// VideoMaterialClearTaskStatus_DONE 已完成
	VideoMaterialClearTaskStatus_DONE VideoMaterialClearTaskStatus = "DONE"
	// VideoMaterialClearTaskStatus_RUNNING 运行中
	VideoMaterialClearTaskStatus_RUNNING VideoMaterialClearTaskStatus = "RUNNING"
	// VideoMaterialClearTaskStatus_TIMEOUT 超时
	VideoMaterialClearTaskStatus_TIMEOUT VideoMaterialClearTaskStatus = "TIMEOUT"
)
