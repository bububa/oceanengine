package appmanagement

import "github.com/bububa/oceanengine/marketing-api/enum"

// File 文件内容
type File struct {
	// UploadID 文件上传id
	UploadID uint64 `json:"upload_id,omitempty"`
	// FileType 文件类型，允许值：
	// APK APK包文件、IMAGE 图片文件、VIDEO 视频文件
	FileType string `json:"file_type,omitempty"`
	// FileTag 文件标示，允许值
	// DEFAULT 默认、AGE_REMINDER 适龄提醒、ANTI_ADDICTION_TIPS 防沉迷提示、MATERIAL_SCREENSHOT 素材截图、REAL_NAME_VERIFIED 实名注册
	FileTag enum.AppFileTag `json:"file_tag,omitempty"`
}
