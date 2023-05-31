package wechat

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// InstanceListRequest 获取微信号码包列表 API Request
type InstanceListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filter 过滤条件
	Filter *InstanceListFilter `json:"filter,omitempty"`
}

// InstanceListFilter 过滤条件
type InstanceListFilter struct {
	// InstanceIDs 微信号码包IDs，超过100个时报参数错误
	InstanceIDs []uint64 `json:"instance_ids,omitempty"`
	// Name 微信号码包名称，模糊匹配
	Name string `json:"name,omitempty"`
	// WechatNumber 微信号，精确匹配
	WechatNumber string `json:"wechat_number,omitempty"`
	// CreateTimeStart 创建时间起始日期，格式：2022-07-19
	CreateTimeStart string `json:"create_time_start,omitempty"`
	// CreateTimeEnd 创建时间截止日期，格式：2022-07-29
	CreateTimeEnd string `json:"create_time_end,omitempty"`
	// Page 分页，>=1，默认值1
	Page int `json:"page,omitempty"`
	// PageSize 页大小，1-100，默认值20
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r InstanceListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filter != nil {
		values.Set("filter", string(util.JSONMarshal(r.Filter)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// InstanceListResponse 获取微信号码包列表 API Response
type InstanceListResponse struct {
	model.BaseResponse
	Data *InstanceListData `json:"data,omitempty"`
}

type InstanceListData struct {
	// TotalCount 总数
	TotalCount int `json:"total_count,omitempty"`
	// TotalPage 总页数
	TotalPage int `json:"total_page,omitempty"`
	// Items 微信号码包列表返回数据
	Items []Instance `json:"items,omitempty"`
}
