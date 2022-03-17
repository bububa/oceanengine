package file

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoUpdateRequest 更新视频API Request
type VideoUpdateRequest struct {
	// AdvertiserID 素材归属的广告主
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Videos 限制最多 100 个
	Videos []VideoForUpdate `json:"videos,omitempty"`
}

// VideoForUpdate 上传视频信息
type VideoForUpdate struct {
	// VideoID 视频id，更新必填
	VideoID string `json:"video_id,omitempty"`
	// Filename 视频素材名称，最多255个字，更新时必填
	Filename string `json:"filename,omitempty"`
	// StatusCode 更新状态
	// SUCCESS 成功; FAILURE 失败（原因：操作视频非当前广告主所有或操作的视频不存在）
	StatusCode enum.VideoUpdateStatusCode `json:"status_code,omitempty"`
}

// Encode implement PostRequest interface
func (r VideoUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// VideoUpdateResponse 更新视频API Response
type VideoUpdateResponse struct {
	model.BaseResponse
	Data *VideoUpdateResponseData `json:"data,omitempty"`
}

// VideoUpdateResponseData json返回值
type VideoUpdateResponseData struct {
	// Videos 视频列表
	Videos []VideoForUpdate `json:"videos,omitempty"`
}
