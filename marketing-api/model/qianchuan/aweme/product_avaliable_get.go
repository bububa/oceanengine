package aweme

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/product"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ProductAvailableGetRequest 达人获取可投商品列表 API Request
type ProductAvailableGetRequest struct {
	// 千川广告主账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 需要拉取商品列表的抖音id
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// Cursor 页码游标值，第一次拉取，无需入参
	Cursor int `json:"cursor,omitempty"`
}

// Encode implement GetRequest interface
func (r ProductAvailableGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("aweme_id", strconv.FormatUint(r.AwemeID, 10))
	if r.Cursor > 0 {
		values.Set("cursor", strconv.Itoa(r.Cursor))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

type ProductAvailableGetResponse struct {
	model.BaseResponse
	Data *ProductAvailableGetResult `json:"data,omitempty"`
}

type ProductAvailableGetResult struct {
	// ProductList 商品列表
	ProductList []product.Product `json:"product_list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
