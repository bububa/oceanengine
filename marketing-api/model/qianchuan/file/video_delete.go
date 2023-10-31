package file

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoDeleteRequest 批量删除视频素材 API Request
type VideoDeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// VideoIDs 待删除的video_id列表，长度范围：1 ~ 100
	VideoIDs []string `json:"video_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r VideoDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// VideoDeleteResponse 批量删除视频素材 API Response
type VideoDeleteResponse struct {
	model.BaseResponse
	Data struct {
		// FailVideoIDs 操作失败的video_id列表，不在此列表内的素材表示删除成功
		FailVideoIDs []string `json:"fail_video_ids,omitempty"`
	} `json:"data,omitempty"`
}
