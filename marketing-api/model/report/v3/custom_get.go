package v3

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CustomGetRequest 自定义报表 API Request
type CustomGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// DateTopic 数据主题
	DataTopic string `json:"data_topic"`
	// Dimensions 维度列表。获取方式：巨量引擎体验版—>报表—>新建/编辑自定义报表—>API参数生成。该字段从前端自定义报表中获取，建议不要修改。
	Dimensions []string `json:"dimensions,omitempty"`
	// Metrics 指标列表 。获取方式：巨量引擎体验版—>报表—>新建/编辑自定义报表—>API参数生成。该字段从前端自定义报表中获取，建议不要修改。
	Metrics []string `json:"metrics,omitempty"`
	// StartTime 起始日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期
	StartTime time.Time `json:"start_time,omitempty"`
	// EndTime 结束日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期，时间跨度不能超过30天
	EndTime time.Time `json:"end_time,omitempty"`
	// OrderBy 排序
	OrderBy []OrderBy `json:"order_by,omitempty"`
	// Page 页码；默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，即每页展示的数据量；默认值: 20；取值范围: 1-1000
	PageSize int `json:"page_size,omitempty"`
	// Filters 过滤字段，json格式，支持字段如下
	Filters []CustomGetFilter `json:"filters,omitempty"`
}

// OrderBy 排序
type OrderBy struct {
	// Field 排序字段。字段必须在选中的metrics或dimensions中。其中metrics所有字段支持排序。dimensions是否排序请参考维度、指标说明。
	Field string `json:"field,omitempty"`
	// Type 排序方式；默认值: DESC；允许值: ASC, DESC
	Type enum.OrderType `json:"type,omitempty"`
}

// CustomGetFilter 数据报表过滤条件
type CustomGetFilter struct {
	// Field 过滤的消耗指标字段
	Field string `json:"field,omitempty"`
	// Type 字段类型。允许值：
	// 1 -固定枚举值
	// 2 - 固定输入值
	// 3 -数值类型
	Type int `json:"type,omitempty"`
	// Operator 处理方式。 允许值：
	// 1 -等于
	// 2 -小于
	// 3 -小于等于
	// 4 -大于
	// 5 -大于等于
	// 6 -不等于
	// 7-包含
	// 8 -不包含
	// 9 -范围查询
	// 10 -多个值匹配包含
	// 11 -多个值匹配都要包含
	Operator int `json:"operator,omitempty"`
	// Values 过滤字段具体值
	Values []string `json:"values,omitempty"`
}

// Encode implement GetRequest interface
func (r CustomGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("start_time", r.StartTime.Format("2006-01-02"))
	values.Set("end_time", r.EndTime.Format("2006-01-02"))
	if r.AdvertiserID > 0 {
		values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	}
	if len(r.DataTopic) > 0 {
		values.Set("data_topic", r.DataTopic)
	}
	if len(r.Dimensions) > 0 {
		fields, _ := json.Marshal(r.Dimensions)
		values.Set("dimensions", string(fields))
	}
	if len(r.Metrics) > 0 {
		fields, _ := json.Marshal(r.Metrics)
		values.Set("metrics", string(fields))
	}
	if len(r.OrderBy) > 0 {
		fields, _ := json.Marshal(r.OrderBy)
		values.Set("order_by", string(fields))
	} else {
		values.Set("order_by", "[]")
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if len(r.Filters) > 0 {
		filters, _ := json.Marshal(r.Filters)
		values.Set("filters", string(filters))
	} else {
		values.Set("filters", "[]")
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// CustomGetResponse 自定义数据报表 API Response
type CustomGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CustomGetResult `json:"data,omitempty"`
}

// CustomGetResult 返回值
type CustomGetResult struct {
	// Rows 数据列表
	Rows []CustomGetListItem `json:"rows,omitempty"`
	// TotalMetrics 指标汇总数据
	TotalMetrics *CustomMetrics `json:"total_metrics,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// CustomGetListItem 数据详情
type CustomGetListItem struct {
	// Metrics 指标数据
	Metrics *CustomMetrics `json:"metrics,omitempty"`
	// Dimensions 维度数据
	Dimensions *CustomDimensions `json:"dimensions,omitempty"`
}
