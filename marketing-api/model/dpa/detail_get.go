package dpa

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DetailGetRequest 获取商品列表 API Request
type DetailGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ProductPlatformID 商品库ID ，可通过【获取商品库信息】获取
	ProductPlatformID uint64 `json:"product_platform_id,omitempty"`
	// Filtering 过滤条件
	Filtering *DetailGetFilter `json:"filtering,omitempty"`
	// Page 页码， 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量， 默认值: 20，最大 100
	PageSize int `json:"page_size,omitempty"`
}

// DetailGetFilter 过滤条件
type DetailGetFilter struct {
	// DpaCategories 商品行业分类，长度限制 50，使用的三级行业ID ，可从 【获取行业列表】 接口获取
	DpaCategories []uint64 `json:"dpa_categories,omitempty"`
	// IsRecommend 是否系统推荐，1表示只看系统推荐，0表示不限系统推荐，默认为0
	IsRecommend int `json:"is_recommend,omitempty"`
	// ProductID 商品ID，精准匹配，与product_name（商品名称）同时传入，以商品ID为准
	ProductID uint64 `json:"product_id,omitempty"`
	// ProductName 商品名称，模糊匹配，长度限制50字，与product_id（商品ID）同时传入，以商品ID为准
	ProductName string `json:"product_name,omitempty"`
	// Status 商品投放状态，枚举值：0代表不可投放，1代表可投放
	Status *int `json:"status,omitempty"`
}

// Encode implement GetRequest interface
func (r DetailGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("product_platform_id", strconv.FormatUint(r.ProductPlatformID, 10))
	if r.Filtering != nil {
		bs, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(bs))
	}
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// DetailGetResponse 获取商品列表 API Response
type DetailGetResponse struct {
	model.BaseResponse
	// Data 返回数据
	Data *DetailGetResponseData `json:"data,omitempty"`
}

// DetailGetResponseData 返回数据
type DetailGetResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 商品库商品列表
	List []Product `json:"list,omitempty"`
}
