package ad

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// RecommendKeywordsGetRequest 获取系统推荐的搜索关键词 API Request
type RecommendKeywordsGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	// 1. “search_word”、“product_id”、“aweme_id”，必须传入一个；
	// 2. 优先按“search_word”过滤，其次按“product_id”过滤，最后按“aweme_id”过滤
	Filtering *RecommendKeywordsFilter `json:"filtering,omitempty"`
	// OrderField 排序字段：允许值:
	// COMPETITION: 竞争程度
	// DEFAULT: 推荐度（默认）
	// PV: 月搜索量
	// RELEVANCE: 相关性
	OrderField qianchuan.RecommendKeywordOrderField `json:"order_field,omitempty"`
	// OrderType 排序方式，允许值：
	// ASC: 升序
	// DESC: 降序
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// CacheID 选填入参。
	// 请求唯一标识，能保证同样的过滤条件返回的数据是一致的。
	// 使用过期后的cache_id请求该接口，不会报错，会当作全新的查询操作来处理。
	// 注意：每次全新的查询，返回的数据可能会不一样
	CacheID string `json:"cache_id,omitempty"`
	// Page 页码，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，允许值：500, 1000，默认值：500
	PageSize int `json:"page_size,omitempty"`
}

type RecommendKeywordsFilter struct {
	// SearchWord 词根
	SearchWord string `json:"search_word,omitempty"`
	// ProductID 商品ID，此参数用于获取与商品更为相关的关键词推荐
	// 注：PC侧创建「短视频带货」搜索广告时，默认根据product_id获取推荐关键
	ProductID uint64 `json:"product_id,omitempty"`
	// AwemeID 抖音id，与计划创建时设置的抖音ID一致
	// 注：PC侧创建「直播带货」搜索广告时，默认根据aweme_id获取推荐关键
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// SuggestReason 推荐理由类型：允许值:
	// CLICK: 高点击词
	// DARKHORSE: 黑马词
	// LOW_COST: 低成本词
	// POTENTIAL: 蓝海词
	// PREFERENCE: 同行买词
	SuggestReason qianchuan.KeywordSuggestReason `json:"suggest_reason,omitempty"`
}

// Encode implement GetRequest interface
func (r RecommendKeywordsGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.OrderField != "" {
		values.Set("order_field", string(r.OrderField))
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.CacheID != "" {
		values.Set("cache_id", r.CacheID)
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// RecommendKeywordsGetResponse 获取系统推荐的搜索关键词 API Response
type RecommendKeywordsGetResponse struct {
	model.BaseResponse
	Data *RecommendKeywordsGetResult `json:"data,omitempty"`
}

type RecommendKeywordsGetResult struct {
	// List 关键词列表
	List []RecommendKeyword `json:"list,omitempty"`
	// PageInfo
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
