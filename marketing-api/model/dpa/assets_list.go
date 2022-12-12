package dpa

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AssetsListRequest 获取投放条件列表 API Request
type AssetsListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PlatformID 商品库ID
	PlatformID uint64 `json:"platform_id,omitempty"`
	// ProductIDs 商品ID列表
	ProductIDs []uint64 `json:"product_ids,omitempty"`
	// Filtering 过滤条件
	Filtering *AssetsListFilter `json:"filtering,omitempty"`
	// Page 页码， 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量， 默认值: 10，最大 100
	PageSize int `json:"page_size,omitempty"`
}

// AssetsListFilter 过滤条件
type AssetsListFilter struct {
	// Status 物件状态， 0代表暂停， 1代表启用
	Status int `json:"status"`
}

// Encode implement GetRequest interface
func (r AssetsListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("platform_id", strconv.FormatUint(r.PlatformID, 10))
	bs, _ := json.Marshal(r.ProductIDs)
	values.Set("product_ids", string(bs))
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

// AssetsListResponse 获取投放条件列表 API Response
type AssetsListResponse struct {
	model.BaseResponse
	Data *AssetsListResponseData `json:"data,omitempty"`
}

// AssetsListResponseData
type AssetsListResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 商品库商品列表
	List []Asset `json:"list,omitempty"`
}
