package hotmaterialderive

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdoptRequest 采纳爆款裂变视频 API Request
type AdoptRequest struct {
	// AdvertiserID 广告主账户 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// VideoIDs 采纳的视频 ID，有效期为 7 天（从 task 完成开始计算）单次最多 50 个
	VideoIDs []string `json:"video_ids,omitempty"`
}

// Encode implements PostRequest interface
func (r AdoptRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AdoptResponse 采纳爆款裂变视频 API Response
type AdoptResponse struct {
	model.BaseResponse
	Data struct {
		Results []AdoptResult `json:"results,omitempty"`
	} `json:"data,omitempty"`
}

type AdoptResult struct {
	// VideoID 视频 ID
	VideoID string `json:"video_id,omitempty"`
	// StatusCode 采纳状态码，详见 应答 status_code 与 status_message
	StatusCode int `json:"status_code,omitempty"`
	// StatusMessage 采纳状态信息，详见 应答 status_code 与 status_message
	StatusMessage string `json:"status_message,omitempty"`
	// MaterialInfo 素材详细信息，只有 status_code=0 时候返回
	MaterialInfo *MaterialInfo `json:"material_info,omitempty"`
}

// MaterialInfo 素材详细信息
type MaterialInfo struct {
	// VideoID 视频 ID，采纳有效期为 7 天
	VideoID string `json:"video_id,omitempty"`
	// VideoURL 预览链接，有效期为 7 天
	VideoURL string `json:"video_url,omitempty"`
	// Filename 文件名
	Filename string `json:"filename,omitempty"`
	// CoverURI 封面 URI
	CoverURI string `json:"cover_uri,omitempty"`
	// CoverURL 封面链接，有效期为 7 天
	CoverURL string `json:"cover_url,omitempty"`
	// Duration 视频时长
	Duration float64 `json:"duration,omitempty"`
	// MaterialID 素材 ID
	MaterialID uint64 `json:"material_id,omitempty"`
	// ImageMode 视频类型
	ImageMode int `json:"image_mode,omitempty"`
}
