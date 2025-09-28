package liveroom

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AnalysisGetRequest 直播间分析报表 API Request
type AnalysisGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartTime 报表开始时间，格式为%Y-%m-%d %H:%M:%S
	// 默认从7天前开始，仅支持天级别的数据查询，请从0点开始查询，否则返回值为空
	StartTime string `json:"start_time,omitempty"`
	// EndTime 报表结束时间，格式为%Y-%m-%d %H:%M:%S
	// 默认为今天，最大时间范围为62天
	EndTime string `json:"end_time,omitempty"`
	// Filtering 过滤器
	Filtering *AnalysisGetFilter `json:"filtering,omitempty"`
	// OrderField 排序字段
	OrderField string `json:"order_field,omitempty"`
	// OrderType 排序方式，可选值:
	// ASC
	// DESC
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	PageSize int `json:"page_size,omitempty"`
}

// AnalysisGetFilter 过滤器
type AnalysisGetFilter struct {
	// RoomIDs 直播间ID，最多筛选100个直播间ID
	// 返回数据范围为筛选的直播间下广告和自然流量数据，包含直播间的所有账户数据
	RoomIDs []uint64 `json:"room_ids,omitempty"`
	// AwemeIDs 抖音号ID
	// 返回数据范围为筛选的抖音号下全量直播间数据，包括广告和自然流量，包含直播间的所有账户数据
	AwemeIDs []uint64 `json:"aweme_ids,omitempty"`
	// FirstFlowCategory 直播间一级流量来源，可选值:
	// AD 竞价
	// BRAND 品牌
	// DOU DOU+
	// NATURAL 自然流量
	// OTHER 其他
	FirstFlowCategory string `json:"first_flow_category,omitempty"`
}

// Encode implements GetRequest interface
func (r AnalysisGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.StartTime != "" {
		values.Set("start_time", r.StartTime)
	}
	if r.EndTime != "" {
		values.Set("end_time", r.EndTime)
	}
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
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
