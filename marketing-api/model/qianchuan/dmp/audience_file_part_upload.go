package dmp

import (
	"io"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// AudienceFilePartUploadRequest 大文件分片上传 API Request
type AudienceFilePartUploadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// File 文件
	File io.Reader `json:"file,omitempty"`
	// Filename 文件名
	Filename string `json:"filename,omitempty"`
	// PartNum 分片序号，从0开始保持递增1，将按该序号将各分片内容进行合并
	PartNum int `json:"part_num,omitempty"`
	// FileKey 文件唯一标识，用于标识多个分片属于同一个文件
	// 分片序号为0时，该字段无需设置；但在上传后续分片时，该字段必填，同时使用前一分片上传后接口返回的file_key值作为入參
	FileKey string `json:"file_key,omitempty"`
	// IsFinish 是否结束分片上传，允许值：
	// 1：是
	// 0：否
	IsFinished int `json:"is_finished,omitempty"`
}

// Encode implement UploadReqeust interface
func (r AudienceFilePartUploadRequest) Encode() []model.UploadField {
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
	ret = append(ret, model.UploadField{
		Key:   "part_num",
		Value: strconv.Itoa(r.PartNum),
	})
	if r.FileKey != "" {
		ret = append(ret, model.UploadField{
			Key:   "file_key",
			Value: r.FileKey,
		})
	}
	ret = append(ret, model.UploadField{
		Key:   "is_finished",
		Value: strconv.Itoa(r.IsFinished),
	})
	return ret
}

// AudienceFilePartUploadResponse 大文件分片上传API Response
type AudienceFilePartUploadResponse struct {
	model.BaseResponse
	Data struct {
		// FileKey 文件唯一标识，有效期15天
		// 注意：一个文件的多个分片，文件标识相同
		// 用于后续分片上传时的入參
		FileKey string `json:"file_key,omitempty"`
	} `json:"data,omitempty"`
}
