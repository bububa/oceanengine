package carousel

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AwemeGetRequest 获取抖音号下图文 API Request
type AwemeGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 抖音号uid，需拉取图文的抖音号，可通过「千川账户下已授权抖音号」获取
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// Filtering 图文过滤条件
	Filtering *AwemeGetFilter `json:"filtering,omitempty"`
	// Cursor 页码游标值，第一次拉取，无需入参
	Cursor int `json:"cursor,omitempty"`
	// Count 页面大小，默认值30，限制1-50
	Count int `json:"count,omitempty"`
}

type AwemeGetFilter struct {
	// ProductID 商品id
	ProductID uint64 `json:"product_id,omitempty"`
	// AwemeItemURL 图文url链接
	AwemeItemURL string `json:"aweme_item_url,omitempty"`
	// AwemeCarouselID 挂车商品指定的筛选主页图文id
	AwemeCarouseID uint64 `json:"aweme_carousel_id,omitempty"`
}

// Encode implement GetRequest interface
func (r AwemeGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("aweme_id", strconv.FormatUint(r.AwemeID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
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

// AwemeGetResponse 获取抖音号下图文 API Response
type AwemeGetResponse struct {
	model.BaseResponse
	Data *AwemeGetResult `json:"data,omitempty"`
}

type AwemeGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// ImageAggregateList 图文列表
	ImageAggregateList []AwemeCarousel `json:"image_aggregate_list,omitempty"`
}

// AwemeCarousel 抖音号图文
type AwemeCarousel struct {
	// AwemeCarouselID 图文itemid
	AwemeCarouselID string `json:"aweme_carousel_id,omitempty"`
	// Title 图文标题
	Title string `json:"title,omitempty"`
	// Description 图文描述
	Description string `json:"description,omitempty"`
	// Images 图文图片列表
	Images []Image `json:"images,omitempty"`
	// MusicID 音乐ID
	MusicID string `json:"music_id,omitempty"`
	// MusicURL 音乐播放链接
	MusicURL string `json:"music_url,omitempty"`
	// Metrics 数据指标
	Metrics *Metrics `json:"metrics,omitempty"`
}

// Metrics 数据指标
type Metrics struct {
	// ViewCnt 观看数
	ViewCnt int64 `json:"view_cnt,omitempty"`
	// LikeCnt 点赞数
	LikeCnt int64 `json:"like_cnt,omitempty"`
	// ShareCnt 分享数
	ShareCnt int64 `json:"share_cnt,omitempty"`
	// CommentCnt 评论数
	CommentCnt int64 `json:"comment_cnt,omitempty"`
}
