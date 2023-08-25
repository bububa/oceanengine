package file

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MaterialBindRequest 素材推送 API Request
type MaterialBindRequest struct {
	// AdvertiserID 素材归属广告主
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TargetAdvertiserIDs 待推送的广告主，数量限制：<=50
	TargetAdvertiserIDs []uint64 `json:"target_advertiser_ids,omitempty"`
	// VideoIDs 视频ID，数量限制：<=50; 注意：跟image_ids必须二选一、组织共享视频不可推送
	VideoIDs []string `json:"video_ids,omitempty"`
	// ImageIDs 图片ID，数量限制：<=50; 注意：跟video_ids必须二选一
	ImageIDs []string `json:"image_ids,omitempty"`
	// CarouselIDs 图集ID，数量限制：<=50 注意：跟image_ids、video_ids必须三选一
	CarouselIDs []uint64 `json:"carousel_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r MaterialBindRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// MaterialBindResponse 素材推送 API Response
type MaterialBindResponse struct {
	model.BaseResponse
	Data *MaterialBindResponseData `json:"data,omitempty"`
}

// MaterialBindResponseData json返回值
type MaterialBindResponseData struct {
	// FailList 推送失败列表
	FailList []FailedMaterialBind `json:"fail_list,omitempty"`
}

// FailedMaterialBind 推送失败素材
type FailedMaterialBind struct {
	// VideoID 推送失败的视频id
	VideoID string `json:"video_id,omitempty"`
	// ImageID 推送失败的图片id
	ImageID string `json:"image_id,omitempty"`
	// CarouselID 推送失败的图集id
	CarouselID uint64 `json:"carousel_id,omitempty"`
	// TargetAdvertiserID 目标推送广告主id
	TargetAdvertiserID uint64 `json:"target_advertiser_id,omitempty"`
	// FailReason 推送失败原因
	// 枚举值：VIDEO_BINDING_EXISTED（视频已存在）、IMAGE_BINDING_EXISTED（图片已存在）
	FailReason enum.MaterialBindFailReason `json:"fail_reason,omitempty"`
}
