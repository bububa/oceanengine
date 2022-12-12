package shop

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdvertiserListRequest 获取店铺账户关联的广告账户列表
type AdvertiserListRequest struct {
	// ShopID 店铺id
	ShopID uint64 `json:"shop_id,omitempty"`
	// Page 页码.默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量.默认值: 10， 最大值：100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r AdvertiserListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("shop_id", strconv.FormatUint(r.ShopID, 10))
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

// AdvertiserListResponse 获取店铺账户关联的广告账户列表 API Response
type AdvertiserListResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AdvertiserListResponseData `json:"data,omitempty"`
}

// AdvertiserListResponseData json返回值
type AdvertiserListResponseData struct {
	// PageInfo 分页信息
	PageInfo model.PageInfo `json:"page_info,omitempty"`
	// List 千川广告主账户ID列表
	List []uint64 `json:"list,omitempty"`
}
