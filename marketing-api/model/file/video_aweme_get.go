package file

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoAwemeGetRequest 获取抖音号下的视频 API Request
type VideoAwemeGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 需拉取视频的抖音号
	AwemeID string `json:"aweme_id,omitempty"`
	// Filtering 视频过滤条件
	Filtering *VideoAwemeGetFiltering `json:"filtering,omitempty"`
	// Page 页码，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小， 默认值：10，取值范围1-100
	PageSize int `json:"page_size,omitempty"`
	// 页码游标值，第一次拉取，请传入0
	Cursor string `json:"cursor,omitempty"`
	// 页面数据量
	Count int `json:"count,omitempty"`
}

// VideoAwemeGetFiltering 筛选条件
type VideoAwemeGetFiltering struct {
	// AwemeURL 抖音视频链接，可在抖音端上通过【分享】-【复制链接】获取
	AwemeURL string `json:"aweme_url,omitempty"`
}

// Encode implement GetRequest interface
func (r VideoAwemeGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("aweme_id", r.AwemeID)
	if r.Filtering != nil {
		filter, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filter))
	}
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if r.Count > 0 {
		values.Set("count", strconv.Itoa(r.Count))
		values.Set("cursor", r.Cursor)
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// VideoAwemeGetResponse 获取抖音号下的视频 API Response
type VideoAwemeGetResponse struct {
	model.BaseResponse
	Data *VideoAwemeGetResponseData `json:"data,omitempty"`
}

// VideoAwemeGetResponseData json返回值
type VideoAwemeGetResponseData struct {
	// List 视频列表
	List []VideoAweme `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo   *model.PageInfo `json:"page_info,omitempty"`
	CursorInfo CursorInfo      `json:"cursor_info"`
}

type CursorInfo struct {
	HasMore bool   `json:"has_more"`
	Cursor  string `json:"cursor"`
}

type VideoAweme struct {
	AwemePlayUrl  string  `json:"aweme_play_url"`
	Duration      float64 `json:"duration"`
	ImageMode     string  `json:"image_mode"`
	ItemId        uint64  `json:"item_id"`
	Title         string  `json:"title"`
	VideoCoverId  string  `json:"video_cover_id"`
	VideoCoverUrl string  `json:"video_cover_url"`
	VideoId       string  `json:"video_id"`
}
