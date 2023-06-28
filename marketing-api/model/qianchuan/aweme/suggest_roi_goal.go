package aweme

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SuggestRoiGoalRequest 获取随心推ROI建议出价 API Request
type SuggestRoiGoalRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 推广对象的抖音id，可通过【查询可推广抖音号列表】接口获取名下可推广抖音号
	// 视频加热直播间：传递视频作者的抖音uid
	// 直接加热直播间：传递主播的抖音uid
	AwemeID uint64 `json:"aweme_id,omitempty"`
}

// Encode implement GetRequest interface
func (r SuggestRoiGoalRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("aweme_id", strconv.FormatUint(r.AwemeID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret

}

// SuggestRoiGoalResponse 获取随心推ROI建议出价 API Response
type SuggestRoiGoalResponse struct {
	model.BaseResponse
	Data struct {
		// EcpRoiGoal roi建议出价
		EcpRoiGoal float64 `json:"ecp_roi_goal,omitempty"`
	} `json:"data,omitempty"`
}
