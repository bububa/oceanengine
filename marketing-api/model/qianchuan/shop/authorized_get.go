package shop

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AuthorizedGetRequest 获取广告主绑定的店铺列表 API Request
type AuthorizedGetRequest struct {
	// AdvertiserID 千川广告主账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页码, 默认值： 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	// 默认值：10，最大值：100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r AuthorizedGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
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

// AuthorizedGetResponse 获取广告主绑定的店铺列表 API Response
type AuthorizedGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AuthorizedGetResult `json:"data,omitempty"`
}

type AuthorizedGetResult struct {
	// ShopList shop账号列表
	ShopList []Shop `json:"shop_list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
