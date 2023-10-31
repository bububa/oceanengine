package v3

import (
	"time"

	v3 "github.com/bububa/oceanengine/marketing-api/model/report/v3"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 自定义报表—创建异步任务 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TaskName 任务名称。最大长度 25 个字符，不能为空字符
	TaskName string `json:"task_name,omitempty"`
	// Force true/false。是否强制生成新的任务（不复用之前任务的结果），不传默认false
	Force bool `json:"force,omitempty"`
	// DateTopic 数据主题
	DataTopic string `json:"data_topic"`
	// Dimensions 维度列表。获取方式：巨量引擎体验版—>报表—>新建/编辑自定义报表—>API参数生成。该字段从前端自定义报表中获取，建议不要修改。
	Dimensions []string `json:"dimensions,omitempty"`
	// Metrics 指标列表 。获取方式：巨量引擎体验版—>报表—>新建/编辑自定义报表—>API参数生成。该字段从前端自定义报表中获取，建议不要修改。
	Metrics []string `json:"metrics,omitempty"`
	// Filters 过滤字段，json格式，支持字段如下
	Filters []v3.CustomGetFilter `json:"filters,omitempty"`
	// StartTime 起始日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期
	StartTime time.Time `json:"start_time,omitempty"`
	// EndTime 结束日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期，时间跨度不能超过30天
	EndTime time.Time `json:"end_time,omitempty"`
	// OrderBy 排序
	OrderBy []v3.OrderBy `json:"order_by,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
