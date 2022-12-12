package dpa

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ProductDetailGetRequest 获取商品详情 API Request
type ProductDetailGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PlatformID 商品库ID，可通过【获取商品库信息】获取
	// 若仅填写platform_id不填filtering.product_ids则表示查询该商品库下所有商品
	PlatformID uint64 `json:"platform_id,omitempty"`
	// Filtering 过滤条件
	Filtering *ProductDetailGetFilter `json:"filtering,omitempty"`
	// Page 页码， 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量， 默认值: 20，最大 100
	PageSize int `json:"page_size,omitempty"`
}

// ProductDetailGetFilter 过滤条件
type ProductDetailGetFilter struct {
	// ProductIDs 商品ID
	ProductIDs []uint64 `json:"product_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r ProductDetailGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("platform_id", strconv.FormatUint(r.PlatformID, 10))
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

// ProductDetailGetResponse 获取商品详情 API Response
type ProductDetailGetResponse struct {
	model.BaseResponse
	// Data 返回数据
	Data *DetailGetResponseData `json:"data,omitempty"`
}
