package appmanagement

import "github.com/bububa/oceanengine/marketing-api/enum"

// UploadTask 异步上传任务
type UploadTask struct {
	// UploadID 上传文件id
	UploadID uint64 `json:"upload_id,omitempty"`
	// Status 任务状态，枚举值：
	// FAILED 失败、RUNNING 执行中 、SUCCESS 成功、WAITING 待执行
	Status enum.AppUploadTaskStatus `json:"status,omitempty"`
	// Message 任务信息
	Message string `json:"message,omitempty"`
}
