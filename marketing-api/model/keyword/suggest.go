package keyword

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SuggestRequest 搜索快投关键词推荐 API Request
type SuggestRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 广告计划id，此参数用于获取与当前计划更为相关的关键词推荐
	AdID uint64 `json:"ad_id,omitempty"`
}

// Encode implement GetRequest interface
func (r SuggestRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("ad_id", strconv.FormatUint(r.AdID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// SuggestResponse 搜索快投关键词推荐 API Response
type SuggestResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *SuggestResponseData `json:"data,omitempty"`
}

// SuggestResponseData json返回值
type SuggestResponseData struct {
	// List 推荐关键词列表
	List []SuggestKeyword `json:"list,omitempty"`
}

// SuggestKeyword 推荐关键词
type SuggestKeyword struct {
	// Keyword 推荐关键词
	Keyword string `json:"keyword,omitempty"`
	// MatchType 匹配类型
	MatchType enum.KeywordMatchType `json:"match_type,omitempty"`
	// Msv 月搜索量
	Msv int64 `json:"msv,omitempty"`
}
