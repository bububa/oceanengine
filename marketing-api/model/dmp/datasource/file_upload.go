package datasource

import (
	"io"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// FileUploadRequest 数据源文件上传 API Request
type FileUploadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// File 文件
	File io.Reader `json:"file,omitempty"`
	// Filename 文件名
	Filename string `json:"filename,omitempty"`
	// FileSignature 文件MD5
	FileSignature string `json:"file_signature,omitempty"`
}

// Encode implement UploadReqeust interface
func (r FileUploadRequest) Encode() []model.UploadField {
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
		Key:   "file_signature",
		Value: r.FileSignature,
	})
	return ret
}

// FileUploadResponse 数据文件上传API Response
type FileUploadResponse struct {
	model.BaseResponse
	Data *FileUploadResponseData `json:"data,omitempty"`
}

// FileUploadResponseData json返回值
type FileUploadResponseData struct {
	// FilePath 文件路径，包含作为文件唯一标识的字符串(14天后文件路径过期)
	FilePath string `json:"file_path,omitempty"`
}
