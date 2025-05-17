package star

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ReportDataTopicConfigRequest 获取任务下累计可查询的数据指标 API Request
type ReportDataTopicConfigRequest struct {
	// Topics 数据主题:
	// BASIC_DATA：基础信息
	// FLOW_DATA：流量表现
	// CONVERT_DATA：转化表现
	// SEARCH_DATA：搜索表现
	// RECOMMEND_DATA： 种草表现
	// DY_SHOP_DATA：抖音进店
	// USER_DISTRIBUTION_DATA：用户画像、
	// INDEX_SCORE_DATA： 指数得分
	// COMMENT_DATA：评论数据
	// 直播用户画像仅保留近90天且直播时长 >= 25 分钟直播数据
	Topics []enum.StarReportTopic `json:"topics,omitempty"`
	// StarID 发起请求的客户的starId
	StarID uint64 `json:"star_id,omitempty"`
	// WorkID 交付作品Id，如：视频Id，直播间room_id.
	WorkID uint64 `json:"work_id,omitempty"`
	// DemandID 任务Id
	DemandID uint64 `json:"demand_id,omitempty"`
}

// Encode implement GetRequest interface
func (r ReportDataTopicConfigRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("star_id", strconv.FormatUint(r.StarID, 10))
	values.Set("work_id", strconv.FormatUint(r.WorkID, 10))
	values.Set("demand_id", strconv.FormatUint(r.DemandID, 10))
	values.Set("topics", string(util.JSONMarshal(r.Topics)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ReportDataTopicConfigResponse 获取任务下累计可查询的数据指标 API Response
type ReportDataTopicConfigResponse struct {
	Data *ReportDataTopicConfigResult `json:"data,omitempty"`
	model.BaseResponse
}

type ReportDataTopicConfigResult struct {
	// Stat 对应请求的数据主题数组
	Stat []DataTopicConfig `json:"stat,omitempty"`
}

// DataTopicConfig 任务下累计可查询的数据指标
type DataTopicConfig struct {
	// DataTopic 同请求中的数据主题枚举值。不存在的数据主题将缺省。 可选值:
	// BASIC_DATA 基础信息
	// COMMENT_DATA 评论数据
	// CONVERT_DATA 转化表现
	// DY_SHOP_DATA 抖音进店
	// FLOW_DATA 流量表现
	// INDEX_SCORE_DATA 指数得分
	// RECOMMEND_DATA 种草表现
	// SEARCH_DATA 搜索表现
	// USER_DISTRIBUTION_DATA 用户画像
	DataTopic enum.StarReportTopic `json:"data_topic,omitempty"`
	// Metrics 数据主题下的数据指标，一个数据主题对应一个或多个数据指标。不存在的数据指标将缺省。
	Metrics []DataTopicMetric `json:"metrics,omitempty"`
}

// DataTopicMetric 数据主题下的数据指标
type DataTopicMetric struct {
	// Field 指标字段，仅包含英文与下划线等符号
	Field string `json:"field,omitempty"`
	// Name 指标名称，中文或英文描述字段名称
	Name string `json:"name,omitempty"`
	// Type 指标数据类型：
	// STRING
	// INT64
	// FLOAT64
	// JSON 可选值:
	// FLOAT64 64位浮点数
	// INT64 64位整数
	// JSON JSON
	// STRING 字符串
	Type string `json:"type,omitempty"`
	// Description 指标描述
	Description string `json:"description,omitempty"`
	// Value 指标值
	Value string `json:"value,omitempty"`
}
