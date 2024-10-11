package local

// UploadTaskStatus 任务处理状态
type UploadTaskStatus string

const (
	// UploadTaskStatus_PROCESS 处理中
	UploadTaskStatus_PROCESS UploadTaskStatus = "PROCESS"
	// UploadTaskStatus_SUCCESS 成功
	UploadTaskStatus_SUCCESS UploadTaskStatus = "SUCCESS"
	// UploadTaskStatus_FAILED 失败
	UploadTaskStatus_FAILED UploadTaskStatus = "FAILED"
)
