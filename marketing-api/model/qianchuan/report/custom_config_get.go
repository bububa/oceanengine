package report

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CustomConfigGetRequest 获取自定义报表可用指标和维度 API Request
type CustomConfigGetRequest struct {
	// DataTopics 数据主题列表 可选值：
	// ECP_BASIC_DATA千川平台广告基础数据
	// SITE_PROMOTION_POST_DATA_LIVE  全域推广-素材-直播间画面
	// SITE_PROMOTION_POST_DATA_VIDEO  全域推广-素材-视频
	// SITE_PROMOTION_POST_DATA_OTHER  全域推广-素材-其他创意
	// SITE_PROMOTION_POST_DATA_TITLE  全域推广-素材-标题
	// SITE_PROMOTION_PRODUCT_POST_DATA_IMAGE 商品全域推广-素材-图片
	// SITE_PROMOTION_PRODUCT_POST_DATA_VIDEO 商品全域推广-素材-视频
	// SITE_PROMOTION_PRODUCT_POST_DATA_OTHER  商品全域推广-素材-其他创意
	// SITE_PROMOTION_PRODUCT_POST_DATA_TITLE  商品全域推广-素材-标题
	DataTopics []qianchuan.DataTopic `json:"data_topics,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r CustomConfigGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("data_topics", string(util.JSONMarshal(r.DataTopics)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// CustomConfigGetResponse 获取自定义报表可用指标和维度 API Response
type CustomConfigGetResponse struct {
	Data *CustomConfigGetResult `json:"data,omitempty"`
	model.BaseResponse
}

type CustomConfigGetResult struct {
	// CustomConfigDatas 数据主题配置列表
	CustomConfigDatas []CustomConfigData `json:"custom_config_datas,omitempty"`
}

// CustomConfigData 数据主题配置列表
type CustomConfigData struct {
	// DataTopic 数据主题
	DataTopic string `json:"data_topic,omitempty"`
	// Dimensions 维度列表
	Dimensions []CustomConfigDimension `json:"dimensions,omitempty"`
	// Metrics 指标列表
	Metrics []CustomConfigMetric `json:"metrics,omitempty"`
}

// CustomConfigDimension 维度
type CustomConfigDimension struct {
	// FilterConfig 过滤条件
	FilterConfig *CustomConfigFilterConfig `json:"filter_config,omitempty"`
	// Field 维度字段
	Field string `json:"field,omitempty"`
	// Name 维度名称
	Name string `json:"name,omitempty"`
	// Description 维度描述
	Description string `json:"description,omitempty"`
	// ExclusionDims 互斥的维度列表
	ExclusionDims []string `json:"exclusion_dims,omitempty"`
	// ExclusionMetrics 互斥的指标列表
	ExclusionMetrics []string `json:"exclusion_metrics,omitempty"`
	// Sortable 是否支持排序
	Sortable bool `json:"sortable,omitempty"`
	// Filterable 是否支持筛选
	Filterable bool `json:"filterable,omitempty"`
}

// CustomConfigFilterConfig 过滤条件
type CustomConfigFilterConfig struct {
	// ValueLimit 过滤字段传入数量上限
	ValueLimit json.Number `json:"value_limit,omitempty"`
	// RangeValues 维度指标过滤枚举值列表
	RangeValues []CustomConfigRangeValue `json:"range_values,omitempty"`
	// Type 字段类型
	// 1 -固定枚举值
	// 2 - 固定输入值
	// 3 -数值类型
	Type int `json:"type,omitempty"`
	// Operator 处理方式
	// 7-包含
	Operator int `json:"operator,omitempty"`
}

// CustomConfigRangeValue 维度指标过滤枚举值
type CustomConfigRangeValue struct {
	// Label 维度过滤字段名称
	Label string `json:"label,omitempty"`
	// Value 维度过滤字段具体值
	Value string `json:"value,omitempty"`
}

// CustomConfigMetric 指标
type CustomConfigMetric struct {
	// Field 指标字段
	Field string `json:"field,omitempty"`
	// Name 指标名称
	Name string `json:"name,omitempty"`
	// Unit 指标单位
	Unit json.Number `json:"unit,omitempty"`
	// Description 指标描述
	Description string `json:"description,omitempty"`
	// ExclusionDims 互斥的维度列表
	ExclusionDims []string `json:"exclusion_dims,omitempty"`
}
