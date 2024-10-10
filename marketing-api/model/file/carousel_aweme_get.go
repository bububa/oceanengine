package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CarouselAwemeGetRequest 获取创编可用的抖音图文素材 API Request
type CarouselAwemeGetRequest struct {
	// AdvertiserID 巨量广告广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 抖音号id
	AwemeID string `json:"aweme_id,omitempty"`
	// Filtering 过滤器
	Filtering *CarouselAwemeGetFilter `json:"filtering,omitempty"`
	// Cursor 页码游标值，第一次拉取，请传入0
	// 不传入或者传空字符串默认使用
	Cursor string `json:"cursor,omitempty"`
	// Count 页面数据量，默认10，最大值100
	Count int `json:"count,omitempty"`
}

type CarouselAwemeGetFilter struct {
	// ItemIDs 抖音图文唯一ID
	ItemIDs []uint64 `json:"item_ids,omitempty"`
}

// Encode implements GetRequest interface
func (r CarouselAwemeGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("aweme_id", r.AwemeID)
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Cursor != "" {
		values.Set("cursor", r.Cursor)
	}
	if r.Count > 0 {
		values.Set("count", strconv.Itoa(r.Count))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// CarouselAwemeGetResponse 获取创编可用的抖音图文素材 API Response
type CarouselAwemeGetResponse struct {
	model.BaseResponse
	Data *CarouselAwemeGetResult `json:"data,omitempty"`
}

type CarouselAwemeGetResult struct {
	// CursorInfo 分页信息
	CursorInfo *model.CursorInfo `json:"cursor_info,omitempty"`
	// AwemeCarouseList 抖音图文素材列表
	AwemeCarouseList []AwemeCarousel `json:"aweme_carouse_list,omitempty"`
}
