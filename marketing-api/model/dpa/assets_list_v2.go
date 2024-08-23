package dpa

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AssetsListV2Request 获取投放条件列表(线索版) API Request
type AssetsListV2Request struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UniqueProductIDs 新线索商品ID列表，最多传入100个
	UniqueProductIDs []uint64 `json:"unique_product_ids,omitempty"`
	// Filtering 过滤条件
	Filtering *AssetsListV2Filter `json:"filtering,omitempty"`
	// Page 页码， 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量， 默认值: 10，最大 100
	PageSize int `json:"page_size,omitempty"`
}

// AssetsListV2Filter 过滤条件
type AssetsListV2Filter struct {
	// Status 物件状态，
	// DISABLE 暂停
	// ENABLE 启用
	Status string `json:"status"`
}

// Encode implement GetRequest interface
func (r AssetsListV2Request) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("unique_product_ids", string(util.JSONMarshal(r.UniqueProductIDs)))
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

// AssetsListV2Response 获取投放条件列表 API Response
type AssetsListV2Response struct {
	Data *AssetsListV2Result `json:"data,omitempty"`
	model.BaseResponse
}

// AssetsListV2Result
type AssetsListV2Result struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// AssetList 商品库商品列表
	AssetList []Asset `json:"asset_list,omitempty"`
}
