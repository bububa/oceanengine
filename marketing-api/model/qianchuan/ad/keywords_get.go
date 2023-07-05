package ad

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// KeywordsGetRequest 获取计划的搜索关键词 API Request
type KeywordsGetRequest struct {
	// AdvertiserID 千川广告账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	Filtering *KeywordsGetFilter `json:"filtering,omitempty"`
	// Page 页码，从 1 开始，默认1
	Page int `json:"page,omitempty"`
	// PageSize 每页数目，上限500，默认100
	PageSize int `json:"page_size,omitempty"`
}

type KeywordsGetFilter struct {
	// AdID 待获取搜索关键词的计划ID
	AdID uint64 `json:"ad_id,omitempty"`
}

// Encode implement GetRequest interface
func (r KeywordsGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
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

// KeywordsGetResponse 获取计划的搜索关键词 API Response
type KeywordsGetResponse struct {
	model.BaseResponse
	Data *KeywordsGetResult `json:"data,omitempty"`
}

type KeywordsGetResult struct {
	// List 关键词列表
	List     []Keyword       `json:"list,omitempty"`
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
