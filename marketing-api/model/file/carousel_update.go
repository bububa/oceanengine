package file

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CarouselUpdateRequest 批量更新图集 API Request
type CarouselUpdateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Carousels 需要更新的图集信息
	Carousels []CarouselUpdateItem `json:"carousels,omitempty"`
}

type CarouselUpdateItem struct {
	// ID 图集id
	ID uint64 `json:"id,omitempty"`
	// FileName 更新文件名
	FileName string `json:"file_name,omitempty"`
	// Images 图片主题
	Images []CarouselUpdateImage `json:"images,omitempty"`
}

type CarouselUpdateImage struct {
	// ImageID 图片id
	ImageID string `json:"image_id,omitempty"`
	ImageSubject
}

// Encode implement PostRequest interface
func (r CarouselUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CarouselUpdateResponse 批量更新图集 API Response
type CarouselUpdateResponse struct {
	model.BaseResponse
	Data struct {
		// Carousels 图集信息
		Carousels []CarouselUpdateResult `json:"carousels,omitempty"`
	} `json:"data,omitempty"`
}

type CarouselUpdateResult struct {
	// ID 图片id
	ID uint64 `json:"id,omitempty"`
	// StatusCode 状态码
	StatusCode int `json:"status_code,omitempty"`
	// Message 返回信息
	Message string `json:"message,omitempty"`
}
