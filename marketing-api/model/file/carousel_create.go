package file

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CarouselCreateRequest 上传图集 API Request
type CarouselCreateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Images 图片信息，包含图片id和图片主题
	Images []CarouselImage `json:"images,omitempty"`
	// VideoID 音频id，根据上传视频接口获取
	VideoID string `json:"video_id,omitempty"`
	// FileName 文件名，可自定义素材名，不传则默认取文件名，最长255个字符
	FileName string `json:"file_name,omitempty"`
	// CarouselType 可选值:
	// INFORMATION_FLOW_IMAGE 信息流图集
	// SEARCH_DISPLAY_WINDOW_IMAGE 搜索橱窗橱窗
	// TOUTIAO_SEARCH_AD_IMAGE 搜索大图图集
	CarouselType enum.ImageMode `json:"carousel_type,omitempty"`
}

// Encode implement PostRequest interface
func (r CarouselCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CarouselCreateResponse 上传图集 API Response
type CarouselCreateResponse struct {
	model.BaseResponse
	Data struct {
		// Carousel 图集信息
		Carousel *Carousel `json:"carousel,omitempty"`
	} `json:"data,omitempty"`
}
