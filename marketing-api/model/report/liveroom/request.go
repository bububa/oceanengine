package liveroom

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// Request 直播分析数据报表 API Request
type Request struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartTime 报表开始时间，格式为%Y-%m-%d %H:%M:%S，默认从7天前开始
	StartTime time.Time `json:"start_time,omitempty"`
	// EndTime 报表结束时间，格式为%Y-%m-%d %H:%M:%S，默认为今天，最大时间范围30天
	EndTime time.Time `json:"end_time,omitempty"`
	// Fields 指标字段，默认传入指标"click_product_count", "live_orders_count", "pay_order_gmv", "pay_order_count", "order_convert_rate"
	Fields []string `json:"fields,omitempty"`
	// Filtering 筛选字段
	Filtering *Filtering `json:"filtering,omitempty"`
	// OrderField 排序指标字段，不传默认填空
	OrderField string `json:"order_field,omitempty"`
	// OrderType 排序类型；默认值: DESC；允许值: ASC, DESC，不传入order_field时不保证顺序
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// Page 页码，默认1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认20，不超过100
	PageSize int `json:"page_size,omitempty"`
}

// Filtering 筛选字段
type Filtering struct {
	// ProductIDs 商品id列表，最多传入100
	ProductIDs []uint64 `json:"product_ids,omitempty"`
	// RoomIDs 直播间id列表，最多传入100
	RoomIDs []uint64 `json:"room_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r Request) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if !r.StartTime.IsZero() {
		values.Set("start_time", r.StartTime.Format("2006-01-02 15:04:05"))
	}
	if !r.EndTime.IsZero() {
		values.Set("end_time", r.EndTime.Format("2006-01-02 15:04:05"))
	}
	if r.Fields != nil {
		fields, _ := json.Marshal(r.Fields)
		values.Set("fields", string(fields))
	}
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
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
