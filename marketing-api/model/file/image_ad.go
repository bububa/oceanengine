package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// ImageAdRequest 上传广告图片 API Request
type ImageAdRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UploadType 图片上传方式; 默认值：UPLOAD_BY_FILE; 允许值：UPLOAD_BY_FILE、UPLOAD_BY_URL
	UploadType enum.UploadType `json:"upload_type,omitempty"`
	// ImageSignature 图片的md5值(用于服务端校验)，upload_type为UPLOAD_BY_FILE必填
	ImageSignature string `json:"image_signature,omitempty"`
	// ImageFile 图片文件,upload_type为UPLOAD_BY_FILE必填; 格式: jpg、jpeg、png、bmp、gif, 大小1.5M内
	ImageFile *model.UploadField `json:"image_file,omitempty"`
	// ImageUrl 图片url地址，如http://xxx.xxx，upload_type为UPLOAD_BY_URL必填
	ImageUrl string `json:"image_url,omitempty"`
	// Filename 素材的文件名，可自定义素材名，不传则默认取文件名，最长255个字符; 注：若同一素材已进行上传，重新上传不会改名
	Filename string `json:"filename,omitempty"`
	// IsAigc 图片素材是否为AIGC生成
	IsAigc bool `json:"is_aigc,omitempty"`
}

// Encode implement UploadRequest interface
func (r ImageAdRequest) Encode() []model.UploadField {
	fields := []model.UploadField{
		{
			Key:   "advertiser_id",
			Value: strconv.FormatUint(r.AdvertiserID, 10),
		},
	}
	if r.UploadType == enum.UPLOAD_BY_URL {
		fields = append(fields, model.UploadField{
			Key:   "upload_type",
			Value: string(r.UploadType),
		}, model.UploadField{
			Key:   "image_url",
			Value: r.ImageUrl,
		})
	} else if r.ImageFile != nil {
		filename := r.ImageFile.Value
		if filename == "" {
			filename = "file"
		}
		fields = append(fields, model.UploadField{
			Key:   "image_signature",
			Value: r.ImageSignature,
		}, model.UploadField{
			Key:    "image_file",
			Value:  filename,
			Reader: r.ImageFile.Reader,
		})
	}
	if r.Filename != "" {
		fields = append(fields, model.UploadField{
			Key:   "filename",
			Value: r.Filename,
		})
	}
	if r.IsAigc {
		fields = append(fields, model.UploadField{
			Key:   "is_aigc",
			Value: "true",
		})
	}
	return fields
}

// ImageAdResponse 上传广告图片 API Response
type ImageAdResponse struct {
	model.BaseResponse
	Data *Image `json:"data,omitempty"`
}
