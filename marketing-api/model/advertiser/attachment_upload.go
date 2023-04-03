package advertiser

import (
	"io"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// AttachmentUploadRequest 批量上传资质附件 API Request
type AttachmentUploadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ImageData 图片数据
	// 上传格式支持jpg、 jpeg、 png、 webp、 awebp、heic、gif、tiff、ico、svg、bmp，大小10M内
	ImageData io.Reader `json:"image_data,omitempty"`
	// Filename 文件名
	Filename string `json:"filename,omitempty"`
	// AttachmentType 可选值:
	// DELIVERY_ATTACHMENT: 投放资质附件
	AttachmentType enum.AttachmentType `json:"attachment_type,omitempty"`
}

// Encode implement UploadReqeust interface
func (r AttachmentUploadRequest) Encode() []model.UploadField {
	ret := make([]model.UploadField, 0, 4)
	ret = append(ret, model.UploadField{
		Key:   "advertiser_id",
		Value: strconv.FormatUint(r.AdvertiserID, 10),
	})
	ret = append(ret, model.UploadField{
		Key:    "image_data",
		Value:  r.Filename,
		Reader: r.ImageData,
	})
	ret = append(ret, model.UploadField{
		Key:   "filename",
		Value: r.Filename,
	})
	ret = append(ret, model.UploadField{
		Key:   "attachment_type",
		Value: string(r.AttachmentType),
	})
	return ret
}

// AttachmentUploadResponse 批量上传资质附件 API Response
type AttachmentUploadResponse struct {
	model.BaseResponse
	Data struct {
		// AttachmentID 附件id
		AttachmentID uint64 `json:"attachment_id,omitempty"`
	} `json:"data,omitempty"`
}
