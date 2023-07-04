package appmanagement

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UploadTaskCreateRequest 创建异步文件上传任务 API Request
type UploadTaskCreateRequest struct {
	// AccountID 账户id指可以接的账号体系如广告主id、巨量纵横组织id等
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 账户类型，允许值:
	// AD 广告主账户、BP 巨量纵横组织账号
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// DownloadURL 上传文件下载链接，长度最长255
	DownloadURL string `json:"download_url,omitempty"`
	// Md5 文件md5，用于文件校验
	Md5 string `json:"md5,omitempty"`
	// FileType 文件类型，允许值
	// APK APK包文件、IMAGE 图片文件、VIDEO 视频文件
	FileType string `json:"file_type,omitempty"`
}

// Encode implement PostRequest interface
func (r UploadTaskCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UploadTaskCreateResponse 创建异步文件上传任务 API Response
type UploadTaskCreateResponse struct {
	model.BaseResponse
	Data struct {
		// UploadID 上传文件id
		UploadID uint64 `json:"upload_id,omitempty"`
	} `json:"data,omitempty"`
}
