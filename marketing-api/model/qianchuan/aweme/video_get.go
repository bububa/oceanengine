package aweme

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoGetRequest 获取随心推可投视频列表 API Request
type VideoGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 需要拉取视频的抖音号uid
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// MarketingGoal 营销目标，可选值:
	// LIVE_PROM_GOODS: 直播间带货
	// VIDEO_PROM_GOODS: 短视频带货
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// Cursor 页码游标值，第一次拉取，无需入参
	Cursor int `json:"cursor,omitempty"`
	// Count 页面大小，限制1-20，默认值：10
	Count int `json:"count,omitempty"`
}

// Encode implement GetRequest interface
func (r VideoGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("aweme_id", strconv.FormatUint(r.AwemeID, 10))
	if r.MarketingGoal != "" {
		values.Set("marketing_goal", string(r.MarketingGoal))
	}
	if r.Cursor > 0 {
		values.Set("cursor", strconv.Itoa(r.Cursor))
	}
	if r.Count > 0 {
		values.Set("count", strconv.Itoa(r.Count))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// VideoGetResponse 获取随心推可投视频列表 API Response
type VideoGetResponse struct {
	model.BaseResponse
	Data *VideoGetResult `json:"data,omitempty"`
}

type VideoGetResult struct {
	PageInfo  *model.PageInfo `json:"page_info,omitempty"`
	VideoList []Video         `json:"video_list,omitempty"`
}
