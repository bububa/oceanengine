package advertiser

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AvatarSubmitRequest 更新广告主账户头像 API Request
type AvatarSubmitRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ImageID 账户头像id，例：web.business.image/201910225d0d5a39ae2e246645b486
	// 您可调用「获取广告主账户头像ID」接口获取头像的image_id
	// 格式：jpg/jpeg/png/bmp
	// 大小：<=5M
	// 像素：<=300*300
	ImageID string `json:"image_id,omitempty"`
	// SourceInfo 品牌名称
	SourceInfo string `json:"source_info,omitempty"`
}

// Encode implement PostRequest interface
func (r AvatarSubmitRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AvatarSubmitResponse 更新广告主账户头像 API Response
type AvatarSubmitResponse struct {
	model.BaseResponse
	Data struct {
		// AdvertiserID 更新头像成功的广告主账户id
		AdvertiserID model.Uint64 `json:"advertiser_id,omitempty"`
	} `json:"data,omitempty"`
}
