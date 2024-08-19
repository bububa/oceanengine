package report

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/util"
)

// LongTransferOrderDataGetRequest 获取长周期订单数据 API Request
type LongTransferOrderDataGetRequest struct {
	// DataTopic 数据主题：枚举值
	// LONG_TRANSFER_ORDER长周期转化价值-订单明细
	DataTopic string `json:"data_topic,omitempty"`
	// StartTime 开始时间。格式为：yyyy-MM-dd。例如2022-08-01
	StartTime string `json:"start_time,omitempty"`
	// EndTime 结束时间。格式为：yyyy-MM-dd。例如2022-08-01（默认为当天23:59:59）
	EndTime string `json:"end_time,omitempty"`
	// Dimensions 维度列表。
	// 可通过【长周期订单可用维度和指标】接口获取不同数据主题下的可用维度和指标
	// 注意：lower_search、attribute_time_type、marketing_goal非维度字段
	Dimensions []string `json:"dimensions,omitempty"`
	// Metrics 指标列表 。
	// 可通过【长周期订单可用维度和指标】接口获取不同数据主题下的可用维度和指标
	Metrics []string `json:"metrics,omitempty"`
	// Filters 过滤条件。
	// 可通过【长周期订单可用维度和指标】接口获取不同数据主题下的可用维度和指标
	// 注意：
	// external_action_fork，不支持筛选和排序
	// ad_name，不支持筛选
	// lower_search为过滤字段，非维度字段，仅支持ad_id过滤
	// attribute_time_type为过滤字段，非维度字段
	// marketing_goal为过滤字段，非维度字段
	Filters []CustomGetFilter `json:"filters,omitempty"`
	// OrderBy 排序
	OrderBy []CustomGetOrderBy `json:"order_by,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r LongTransferOrderDataGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("data_topic", r.DataTopic)
	values.Set("dimensions", string(util.JSONMarshal(r.Dimensions)))
	values.Set("metrics", string(util.JSONMarshal(r.Metrics)))
	if len(r.Filters) > 0 {
		values.Set("filters", string(util.JSONMarshal(r.Filters)))
	}
	values.Set("start_time", r.StartTime)
	values.Set("end_time", r.EndTime)
	if len(r.OrderBy) > 0 {
		values.Set("order_by", string(util.JSONMarshal(r.OrderBy)))
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
