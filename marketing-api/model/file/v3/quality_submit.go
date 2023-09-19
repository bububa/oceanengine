package v3

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QualitySubmitRequest 连山视频投前分析提交 API Request
type QualitySubmitRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MaterialChannel 素材业务线 可选值:
	// AD
	// DOU_PLUS
	// QC
	MaterialChannel enum.MaterialChannel `json:"material_channel,omitempty"`
	// MaterialType 素材类型 可选值:
	// VIDEO
	MaterialType string `json:"material_type,omitempty"`
	// VideoURL 仅支持连山素材内网地址：https://***.tos-cn-beijing.ivolces.com/***
	VideoURL string `json:"video_url,omitempty"`
}

// Encode implement PostRequest interface
func (r QualitySubmitRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// QualitySubmitResponse 连山视频投前分析提交 API Response
type QualitySubmitResponse struct {
	model.BaseResponse
	Data struct {
		// MaterialID 预审ID
		MaterialID uint64 `json:"material_id,omitempty"`
	} `json:"data,omitempty"`
}
