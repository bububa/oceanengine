package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoCoverSuggestRequest 获取视频智能封面 API Request
type VideoCoverSuggestRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// VideoID 视频id
	VideoID string `json:"video_id,omitempty"`
}

// Encode implement GetRequest interface
func (r VideoCoverSuggestRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("video_id", r.VideoID)
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// VideoCoverSuggestResponse 获取视频智能封面 API Response
type VideoCoverSuggestResponse struct {
	model.BaseResponse
	Data *VideoCoverSuggestResponseData `json:"data,omitempty"`
}

// VideoCoverSuggestResponseData json返回值
type VideoCoverSuggestResponseData struct {
	// Status 封面生成的状态
	// 枚举值：RUNNING（生成中）、SUCCESS（成功）、FAILED（失败）
	// 注意：视频封面不是实时生成，需要根据status字段判断封面生成的状态，大概需要6，7s中生成，所以当status=RUNNING时请等待后重试
	Status enum.VideoCoverStatus `json:"status,omitempty"`
	// List 视频封面列表
	List []Image `json:"list,omitempty"`
}
