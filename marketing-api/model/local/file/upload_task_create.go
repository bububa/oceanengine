package file

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UploadTaskCreateRequest 异步上传本地推视频 API Request
type UploadTaskCreateRequest struct {
	// LocalAccountID 本地推账户id
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// Filename 素材的文件名
	Filename string `json:"filename,omitempty"`
	// VideoURL 视频 url地址，最大支持上传文件大小：1000M
	// 仅支持开发者购置连山云素材服务上传生成的tos链接上传，不支持其他三方链接地址，具体见接入指南
	VideoURL string `json:"video_url,omitempty"`
}

// Encode implements PostRequest interface
func (r UploadTaskCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UploadTaskCreateResponse 异步上传本地推视频 API Response
type UploadTaskCreateResponse struct {
	model.BaseResponse
	Data struct {
		// TaskID 本地推任务id
		TaskID uint64 `json:"task_id,omitempty"`
	} `json:"data,omitempty"`
}
