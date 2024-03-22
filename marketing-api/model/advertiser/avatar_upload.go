package advertiser

import (
	"io"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// AvatarUploadRequest 获取广告主账户头像ID API Request
type AvatarUploadRequest struct {
	// AdvertiserID 广告主账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ImageData 图片数据
	// 格式要求：bmp,jpeg,jpg,png
	// 文件大小：<=5M
	// 像素：<=300*300
	ImageData io.Reader `json:"image_data,omitempty"`
	// Filename 文件名
	Filename string `json:"filename,omitempty"`
}

// Encode implement UploadReqeust interface
func (r AvatarUploadRequest) Encode() []model.UploadField {
	ret := make([]model.UploadField, 0, 4)
	ret = append(ret, model.UploadField{
		Key:   "advertiser_id",
		Value: strconv.FormatUint(r.AdvertiserID, 10),
	})
	ret = append(ret, model.UploadField{
		Key:    "image_file",
		Value:  r.Filename,
		Reader: r.ImageData,
	})
	return ret
}

// AvatarUploadResponse 获取广告主账户头像ID API Response
type AvatarUploadResponse struct {
	model.BaseResponse
	Data struct {
		// ImageID 账户头像图片ID，您可使用此ID前往「更新广告主账户头像」接口更新头像
		ImageID string `json:"image_id,omitempty"`
	} `json:"data,omitempty"`
}
