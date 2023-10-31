package product

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AvailableGetRequest 获取可投商品列表 API Request
type AvailableGetRequest struct {
	// AdvertiserID 千川广告主账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 抖音号id
	// 注意：在渠道品投放时，需要根据aweme_id来拉取渠道品信息
	AwemeID uint64 `json:"aweme_id,omitempty"`
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
	// MarketingScene 营销场景，可选值，默认FEED：
	// FEED通投
	// SEARCH搜索
	// SHOPPING_MALL 商城广告
	// 注意：如果需要获取商品是否支持新品加速信息，该字段必填，仅允许SEARCH
	MarketingScene qianchuan.MarketingScene `json:"marketing_scene,omitempty"`
}

// Encode implement GetRequest interface
func (r AvailableGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.AwemeID > 0 {
		values.Set("aweme_id", strconv.FormatUint(r.AwemeID, 10))
	}
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
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
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
