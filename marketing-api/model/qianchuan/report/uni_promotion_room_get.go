package report

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UniPromotionDimensionGetRequest 全域推广维度数据 API Request
type UniPromotionDimensionGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 抖音号id
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// RoomID 直播间id
	RoomID uint64 `json:"room_id,omitempty"`
	// Metrics 需要查询的消耗指标，见返回参数，默认返回stat_cost
	Metrics []string `json:"metrics,omitempty"`
	// StartDate 起始日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期，时间跨度不能超过30天
	EndDate string `json:"end_date,omitempty"`
	// Dimension 数据维度 ，如果不传，返回查询日期内的汇总数据
	// 允许值:
	// TIME_GRANULARITY_DAILY (按天维度),会返回每天的数据
	// TIME_GRANULARITY_HOURLY (按小时维度)，会返回每小时维度的数据
	Dimension enum.TimeGranularity `json:"time_granularity,omitempty"`
	// OrderType 排序方式，允许值：
	// ASC 升序（默认）
	// DESC 降序
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// OrderField 排序字段，允许值参考数据指标，见返回字段，默认不传为stat_cost，排序字段需在metrics中
	OrderField string `json:"order_field,omitempty"`
	// Page 页码，  默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，  默认值: 10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r UniPromotionDimensionGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.AwemeID > 0 {
		values.Set("aweme_id", strconv.FormatUint(r.AwemeID, 10))
	}
	if r.RoomID > 0 {
		values.Set("room_id", strconv.FormatUint(r.RoomID, 10))
	}
	if len(r.Metrics) > 0 {
		values.Set("metrics", string(util.JSONMarshal(r.Metrics)))
	}
	if r.Dimension != "" {
		values.Set("time_granularity", string(r.Dimension))
	}
	values.Set("start_date", r.StartDate)
	values.Set("end_date", r.EndDate)
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
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

// UniPromotionDimensionGetResponse 全域推广维度数据 API Response
type UniPromotionDimensionGetResponse struct {
	model.BaseResponse
	Data *UniPromotionDimensionGetResult `json:"data,omitempty"`
}

type UniPromotionDimensionGetResult struct {
	PageInfo *model.PageInfo     `json:"page_info,omitempty"`
	List     []UniPromotionStats `json:"list,omitempty"`
}
