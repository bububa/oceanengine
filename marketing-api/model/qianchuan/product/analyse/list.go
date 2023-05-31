package analyse

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ListRequest 获取商品竞争分析列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Keyword 商品关键词，支持根据商品id/名称搜索
	Keyword string `json:"keyword,omitempty"`
	// TimeRange 商品数据范围，可选值
	// 7：7天（默认）
	// 15：15天
	// 30：30天
	TimeRange int `json:"time_range,omitempty"`
	// OrderType 排序方式，允许值：
	// ASC 升序
	// DESC降序（默认）
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// OrderField 排序字段，允许值参考数据指标，默认不传为stat_cost
	OrderField string `json:"order_field,omitempty"`
	// Page 页码，默认1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认20，不超过100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r ListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Keyword != "" {
		values.Set("keyword", r.Keyword)
	}
	if r.TimeRange > 0 {
		values.Set("time_range", strconv.Itoa(r.TimeRange))
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("paeg_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ListResponse 获取商品竞争分析列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResponseData `json:"data,omitempty"`
}

type ListResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// ProductList 商品列表
	ProductList []Product `json:"product_list,omitempty"`
}
