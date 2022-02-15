package site

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// GetRequest 获取橙子建站站点列表 API Request
type GetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页数默认值: 1，page必须大于0
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值:10，page_size范围为1-1000
	PageSize int `json:"page_size,omitempty"`
	// Status 建站粗粒度状态
	Status enum.SiteSearchStatus `json:"status,omitempty"`
	// Filtering 过滤字段
	Filtering *GetFiltering `json:"filtering,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if r.Status != "" {
		values.Set("status", string(r.Status))
	}
	return values.Encode()
}

// GetFiltering 过滤字段
type GetFiltering struct {
	// Search 搜索字段，按照建站id和建站name进行模糊匹配，范围：1 <= 长度 <= 50
	Search string `json:"search,omitempty"`
}

// GetResponse 获取橙子建站站点列表 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData json返回值
type GetResponseData struct {
	// List 站点列表
	List []Site `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
