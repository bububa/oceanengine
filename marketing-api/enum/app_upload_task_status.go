package enum

// AppUploadTaskStatus 任务状态
type AppUploadTaskStatus string

const (
	// AppUploadTaskStatus_FAILED 失败
	AppUploadTaskStatus_FAILED AppUploadTaskStatus = "FAILED"
	// AppUploadTaskStatus_RUNNING 执行中
	AppUploadTaskStatus_RUNNING AppUploadTaskStatus = "RUNNING"
	// AppUploadTaskStatus_SUCCESS 成功
	AppUploadTaskStatus_SUCCESS AppUploadTaskStatus = "SUCCESS"
	// AppUploadTaskStatus_WAITING 待执行
	AppUploadTaskStatus_WAITING AppUploadTaskStatus = "WAITING"
)
