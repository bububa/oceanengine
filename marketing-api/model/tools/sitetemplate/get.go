package sitetemplate

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取站点模版列表 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filter 过滤条件
	Filter *GetFilter `json:"filter,omitempty"`
	// Page 页码，默认值是1，范围page >=1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值是20，范围1~100
	PageSize int `json:"page_size,omitempty"`
}

// GetFilter 过滤条件
type GetFilter struct {
	// TemplateIDs 模板ID数组，数量限制100
	TemplateIDs []uint64 `json:"template_ids,omitempty"`
	// SiteIDs 站点ID，可通过【橙子建站】平台或【获取橙子建站站点列表】获取
	SiteIDs []uint64 `json:"site_ids,omitempty"`
	// KeywordOfName 模板名称关键词
	KeywordOfName string `json:"keyword_of_name,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filter != nil {
		filter, _ := json.Marshal(r.Filter)
		values.Set("filter", string(filter))
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

// GetResponse 获取站点模版列表 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData json返回值
type GetResponseData struct {
	// List 模板列表
	List []Template `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
