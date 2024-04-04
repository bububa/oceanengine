package report

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CustomGetRequest 自定义报表 API Request
type CustomGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// DataTopics 数据主题，可选值:
	// ECP_BASIC_DATA 千川平台广告基础数据
	DataTopics []string `json:"data_topics,omitempty"`
	// Dimensions 维度列表。
	// 可通过【获取自定义报表可用维度和指标】接口获取不同数据主题下的可用维度和指标
	Dimensions []string `json:"dimensions,omitempty"`
	// Metrics 指标列表 。
	// 可通过【获取自定义报表可用维度和指标】接口获取不同数据主题下的可用维度和指标
	Metrics []string `json:"metrics,omitempty"`
	// Filters 过滤条件。
	// 可通过【获取自定义报表可用维度和指标】接口获取不同数据主题下的可用维度和指标
	Filters []CustomGetFilter `json:"filters,omitempty"`
	// StartTime 开始时间。格式为：2022-08-01 00:00:00
	StartTime string `json:"start_time,omitempty"`
	// EndTime 结束时间。格式为：2022-08-01 23:59:59
	EndTime string `json:"end_time,omitempty"`
	// OrderBy 排序
	OrderBy []CustomGetOrderBy `json:"order_by,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	PageSize int `json:"page_size,omitempty"`
}

type CustomGetFilter struct {
	// Field 过滤的消耗指标字段
	Field string `json:"field,omitempty"`
	// Type 字段类型。允许值：
	// 1 -固定枚举值
	// 2 - 固定输入值
	// 3 -数值类型
	Type int `json:"type,omitempty"`
	// Operator 处理方式。 允许值：
	// 7-包含
	Operator int `json:"operator,omitempty"`
	// Values 过滤字段具体值
	Values []string `json:"values,omitempty"`
}

// CustomGetOrderBy 排序
type CustomGetOrderBy struct {
	// Type 排序类型 ，允许值
	// 1:升序
	// 2:降序
	Type int `json:"type,omitempty"`
	// Field 排序字段。字段必须在选中的metrics或dimensions中。其中metrics所有字段支持排序。dimensions是否排序请参考维度、指标说明。
	Field string `json:"field,omitempty"`
}

// Encode implement GetRequest interface
func (r CustomGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("data_topics", string(util.JSONMarshal(r.DataTopics)))
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

// CustomGetResponse 自定义报表 API Response
type CustomGetResponse struct {
	model.BaseResponse
	Data *CustomGetResult `json:"data,omitempty"`
}

type CustomGetResult struct {
	// Pagination 分页信息
	Pagination *model.PageInfo `json:"pagination,omitempty"`
	// Rows 数据
	Rows []CustomGetRow `json:"rows,omitempty"`
	// TotalMetrics 指标汇总数据
	TotalMetrics map[string]json.Number `json:"total_metrics,omitempty"`
}

type CustomGetRow struct {
	// Metrics 指标数据
	Metrics map[string]json.Number `json:"metrics,omitempty"`
	// Dimensions 维度数据
	Dimensions map[string]string `json:"dimensions,omitempty"`
}
