package eventmanager

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AllAssetsListRequest 获取账户下资产列表（新） API Request
type AllAssetsListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	Filtering *AllAssetsListFilter `json:"filtering,omitempty"`
	// Page 页数，默认值`1`，最大值`999999`
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值`10`，最大值`100`
	PageSize int `json:"page_size,omitempty"`
}

type AllAssetsListFilter struct {
	// AssetIDs 资产id列表，list长度最大`100`
	AssetIDs []uint64 `json:"asset_ids,omitempty"`
	// AssetType 资产类型
	AssetType enum.AssetType `json:"asset_types,omitempty"`
	// ModifyStartTime 按照资产修改时间筛选，开始时间`YYYY-MM-DD`，必须与结束时间搭配传入
	// 开始时间必须≤结束时间
	ModifyStartTime string `json:"modify_start_time,omitempty"`
	// ModifyEndTime 按照资产修改时间筛选，结束时间`YYYY-MM-DD`，必须与开始时间搭配传入
	// 开始时间必须≤结束时间
	ModifyEndTime string `json:"modify_end_time,omitempty"`
}

// Encode implements GetRequest interface
func (r AllAssetsListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
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

// AllAssetsListResponse 获取账户下资产列表（新） API Response
type AllAssetsListResponse struct {
	model.BaseResponse
	Data *AllAssetsListResult `json:"data,omitempty"`
}

type AllAssetsListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 资产列表
	AssetList []AssetBaseInfo `json:"asset_list,omitempty"`
}
