package enum

// VideoCoverStatus 封面生成的状态
type VideoCoverStatus string

const (
	// VideoCoverStatus_RUNNING 生成中
	VideoCoverStatus_RUNNING VideoCoverStatus = "RUNNING"
	// VideoCoverStatus_SUCCESS 成功
	VideoCoverStatus_SUCCESS VideoCoverStatus = "SUCCESS"
	// VideoCoverStatus_FAILED 失败
	VideoCoverStatus_FAILED VideoCoverStatus = "FAILED"
)
