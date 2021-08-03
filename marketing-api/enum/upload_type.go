package enum

// UploadType 上传方式
type UploadType string

const (
	// UPLOAD_BY_FILE 上传文件
	UPLOAD_BY_FILE UploadType = "UPLOAD_BY_FILE"
	// UPLOAD_BY_URL 通过链接上传
	UPLOAD_BY_URL UploadType = "UPLOAD_BY_URL"
)
