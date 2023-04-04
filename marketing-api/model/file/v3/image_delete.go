package v3

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ImageDeleteRequest 批量删除图片素材 API Request
type ImageDeleteRequest struct {
	// AdvertiserID 素材归属的广告主
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ImageIDs 待删除的图片ID列表，长度范围：1 ~ 100
	ImageIDs []string `json:"image_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r ImageDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ImageDeleteResponse 批量删除图片素材 API Response
type ImageDeleteResponse struct {
	model.BaseResponse

	Data struct {
		// FailImageIDs 操作失败的图片ID列表，不在此列表内的素材表示删除成功
		FailImageIDs []string `json:"fail_image_ids,omitempty"`
	} `json:"data,omitempty"`
}
