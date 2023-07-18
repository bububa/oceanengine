package dmp

import (
	"io"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// AudienceFileUploadRequest 小文件直接上传 API Request
type AudienceFileUploadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// File 文件
	File io.Reader `json:"file,omitempty"`
	// Filename 文件名
	Filename string `json:"filename,omitempty"`
}

// Encode implement UploadReqeust interface
func (r AudienceFileUploadRequest) Encode() []model.UploadField {
	ret := make([]model.UploadField, 0, 3)
	ret = append(ret, model.UploadField{
		Key:   "advertiser_id",
		Value: strconv.FormatUint(r.AdvertiserID, 10),
	})
	ret = append(ret, model.UploadField{
		Key:    "file",
		Value:  r.Filename,
		Reader: r.File,
	})
	return ret
}

// AudienceFileUploadResponse 小文件直接上传API Response
type AudienceFileUploadResponse struct {
	model.BaseResponse
	Data *AudienceFileUploadResult `json:"data,omitempty"`
}

// AudienceFileUploadResult json返回值
type AudienceFileUploadResult struct {
	// FileKey 文件唯一标识，有效期15天
	FileKey string `json:"file_key,omitempty"`
	// FileMd5 文件上传成功后返回的md5，可用于校验与本地文件的一致性
	FileMd5 string `json:"file_md5,omitempty"`
}
