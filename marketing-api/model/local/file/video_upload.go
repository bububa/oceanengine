package file

import (
	"io"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// VideoUploadRequest 上传视频 API Request
type VideoUploadRequest struct {
	// LocalAccountID 本地推账户id
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// VideoFile 文件
	VideoFile io.Reader `json:"video_file,omitempty"`
	// Filename 文件名
	Filename string `json:"filename,omitempty"`
	// VideoSignature 文件MD5
	VideoSignature string `json:"video_signature,omitempty"`
}

// Encode implement UploadReqeust interface
func (r VideoUploadRequest) Encode() []model.UploadField {
	ret := make([]model.UploadField, 0, 3)
	ret = append(ret, model.UploadField{
		Key:   "local_account_id",
		Value: strconv.FormatUint(r.LocalAccountID, 10),
	})
	ret = append(ret, model.UploadField{
		Key:    "video_file",
		Value:  r.Filename,
		Reader: r.VideoFile,
	})
	ret = append(ret, model.UploadField{
		Key:   "filename",
		Value: r.Filename,
	})
	ret = append(ret, model.UploadField{
		Key:   "video_signature",
		Value: r.VideoSignature,
	})
	return ret
}

// VideoUploadResponse 上传视频 API Response
type VideoUploadResponse struct {
	model.BaseResponse
	Data *Video `json:"data,omitempty"`
}
