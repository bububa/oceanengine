package product

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/local"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取可投商品列表 API Request
type GetRequest struct {
	// LocalAccountID 本地推广告主ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// LocalDeliveryScene 推广目的，允许值：
	// CONTENT_HEAT 内容加热
	// EXTERNAL 销售线索收集
	// POI_RECOMMEND 门店引流
	// PRODUCT_PAY 团购成交
	LocalDeliveryScene local.LocalDeliveryScene `json:"local_delivery_scene,omitempty"`
	// Filtering 过滤器
	Filtering *GetFilter `json:"filtering,omitempty"`
	// Page 页码，默认值1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值20，最大值100
	PageSize int `json:"page_size,omitempty"`
}

type GetFilter struct {
	// SearchKeyWord 根据商品名称/商品ID搜索（商品名称支持模糊搜索，商品ID支持精确搜索）
	SearchKeyWord string `json:"search_key_word,omitempty"`
}

// Encode implements GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("local_account_id", strconv.FormatUint(r.LocalAccountID, 10))
	if r.LocalDeliveryScene != "" {
		values.Set("local_delivery_scene", string(r.LocalDeliveryScene))
	}
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

// GetResponse 获取可投商品列表 API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Products 商品信息
	Products []Product `json:"products,omitempty"`
}
