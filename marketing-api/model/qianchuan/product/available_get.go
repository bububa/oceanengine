package product

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// AvailableGetRequest 获取可投商品列表 API Request
type AvailableGetRequest struct {
	// AdvertiserID 千川广告主账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤器
	Filtering *GetFiltering `json:"filtering,omitempty"`
	// Page 页码.默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量.默认值: 10， 最大值：100
	PageSize int `json:"page_size,omitempty"`
}

// GetFiltering 过滤器
type GetFiltering struct {
	// ProductIDs 根据商品ID列表筛选，长度：1-100
	ProductIDs []uint64 `json:"product_ids,omitempty"`
	// ProductName 根据商品名称筛选，长度120字符以内，中文算两个字符。
	ProductName string `json:"product_name,omitempty"`
}

// Encode implement GetRequest interface
func (r AvailableGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filter, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filter))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}

// AvailableGetResponse 获取可投商品列表 API Response
type AvailableGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AvailableGetResponseData `json:"data,omitempty"`
}

// AvailalbeGetResponseData json返回值
type AvailableGetResponseData struct {
	// PageInfo 分页信息
	PageInfo model.PageInfo `json:"page_info,omitempty"`
	// ProductList 商品列表
	ProductList []Product `json:"product_list,omitempty"`
}
