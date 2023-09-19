package enterprise

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/enterprise"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ItemListRequest 获取企业号视频列表 API Request
type ItemListRequest struct {
	// CcAccountID 纵横组织ID，纵横组织管理员或协作者授权后，通过【获取已授权账户】接口，查询到“账号角色为CUSTOMER_ADMIN-管理员授权的纵横组织、或CUSTOMER_OPERATOR-协作者授权的纵横组织”的advertiser_id，即为纵横组织ID
	CcAccountID uint64 `json:"cc_account_id"`
	// EDouyinID 企业号账户ID，获取到授权的纵横组织ID后，再通过【获取纵横组织下资产账户列表（分页）】接口，查询到e_douyin_id，即为企业号账户ID，需确保传入的企业号账户ID与纵横组织ID已建立绑定关系，且绑定关系未解除
	EDouyinID string `json:"e_douyin_id"`
	// StartTime 开始时间，默认30天前 只支持天级的。e.g.2022-01-01 开始时间必须大于2021-01-01
	StartTime string `json:"start_time,omitempty"`
	// EndTime 结束时间，默认当天 只支持天级的。e.g.2022-01-02 结束时间必须晚于开始时间，查询时间跨度不能大于3个月
	EndTime string `json:"end_time,omitempty"`
	// Filter 筛选字段
	Filter *ItemListFilter `json:"filter,omitempty"`
	// Page 页数，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，支持范围1-100 之间，默认值：20
	PageSize int `json:"page_size,omitempty"`
}

// ItemListFilter 筛选字段
type ItemListFilter struct {
	// ItemType 视频类型 允许值：ITEM_AD广告视频、ITEM_CONTENT非广告视频(抖音视频)
	ItemType enterprise.ItemType `json:"item_type,omitempty"`
}

// Encode implement GetRequest interface
func (r ItemListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("cc_account_id", strconv.FormatUint(r.CcAccountID, 10))
	values.Set("e_douyin_id", r.EDouyinID)
	if r.StartTime != "" {
		values.Set("start_time", r.StartTime)
	}
	if r.EndTime != "" {
		values.Set("end_time", r.EndTime)
	}
	if r.Filter != nil {
		values.Set("filter", string(util.JSONMarshal(r.Filter)))
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

// ItemListResponse 获取企业号视频列表 API Response
type ItemListResponse struct {
	model.BaseResponse
	Data struct {
		ItemList []Item `json:"item_list,omitempty"`
	} `json:"data,omitempty"`
}
