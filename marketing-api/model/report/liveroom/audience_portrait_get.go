package liveroom

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AudiencePortraitGetRequest 直播间受众分析报表 API Request
type AudiencePortraitGetRequest struct {
	// AdvertiserID 广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Dimension 受众报表维度，可选值:
	// PROVINCE省份
	// AGE年龄
	// PLATFORM平台
	// CITY城市
	// GENDER性别
	Dimension string `json:"dimension,omitempty"`
	// StartTime 开始时间
	StartTime string `json:"start_time,omitempty"`
	// EndTime 结束时间
	EndTime string `json:"end_time,omitempty"`
	// Fields 指标字段
	Fields []string `json:"fields,omitempty"`
	// Filtering 筛选
	Filtering *AudiencePortraitGetFilter `json:"filtering,omitempty"`
	// OrderField 排序指标
	OrderField string `json:"order_field,omitempty"`
	// OrderType 排序类型，可选值:
	// ASC
	// DESC
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	PageSize int `json:"page_size,omitempty"`
}

type AudiencePortraitGetFilter struct {
	// RoomID 直播间ID
	RoomID string `json:"room_id,omitempty"`
}

// Encode implements GetRequest interface
func (r AudiencePortraitGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("dimension", r.Dimension)
	if r.StartTime != "" {
		values.Set("start_time", r.StartTime)
	}
	if r.EndTime != "" {
		values.Set("end_time", r.EndTime)
	}
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
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

// AudiencePortraitGetResponse 直播间受众分析报表 API Response
type AudiencePortraitGetResponse struct {
	model.BaseResponse
	Data *AudiencePortraitGetResult `json:"data,omitempty"`
}

type AudiencePortraitGetResult struct {
	// List 直播间受众分析数据
	List []AudiencePortrait `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// AudiencePortrait 直播间受众分析数据
type AudiencePortrait struct {
	// Age 年龄
	Age string `json:"age,omitempty"`
	// Gender 性别
	Gender string `json:"gender,omitempty"`
	// City 城市
	City string `json:"city,omitempty"`
	// Platform 平台
	Platform string `json:"platform,omitempty"`
	// ProvinceName 省份
	ProvinceName string `json:"province_name,omitempty"`
	// Fields 指标字段
	Fields *Metrics `json:"fields,omitempty"`
}
